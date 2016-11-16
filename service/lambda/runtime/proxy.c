// +build proxy

//
// Copyright 2016 Alsanium, SAS. or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

#include <Python.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

struct handle_return { char* r0; char* r1; };
extern struct handle_return handle(char*, char*, char*);

static PyObject *module_error = NULL;
static PyObject* proxy_log_fn = NULL;
static PyObject* proxy_get_remaining_time_in_millis_fn = NULL;

void
proxy_log(char* msg)
{
  PyObject* tmp = PyObject_CallFunction(proxy_log_fn, "s", msg);
  Py_DECREF(tmp);
  free(msg);
}

long long
proxy_get_remaining_time_in_millis()
{
  PyObject* pms = PyObject_CallFunctionObjArgs(proxy_get_remaining_time_in_millis_fn, NULL);
  unsigned long ms = PyLong_AsLongLong(pms);
  Py_DECREF(pms);
  return ms;
}

static PyObject* py_json_module;
static PyObject* py_os_module;

static void
PyDict_Free(PyObject* p)
{
  PyObject* es = PyDict_Values(p);
  for (Py_ssize_t i = 0; i < PyList_Size(es); i++) {
    PyObject* e =  PyList_GetItem(es, i);
    if (PyDict_Check(e)) {
      PyDict_Free(e);
    } else {
      Py_DECREF(e);
    }
  }
  Py_DECREF(es);
  Py_DECREF(p);
}

static void
PyDict_SetSomeItemString(PyObject* p, char *key, PyObject* val)
{
  if ((val != Py_None) && (!PyDict_Check(val) || (PyDict_Check(val) && PyDict_Size(val)))) {
    PyDict_SetItemString(p, key, val);
  } else {
    Py_DECREF(val);
  }
}

static void
PyDict_CopyItemString(PyObject* p1, PyObject* p2, char *key)
{
  PyDict_SetSomeItemString(p1, key, PyObject_GetAttrString(p2, key));
}

static PyObject*
proxy_marshal_ctx(PyObject* raw)
{
  PyObject* ctx = PyDict_New();

  PyDict_CopyItemString(ctx, raw, "function_name");
  PyDict_CopyItemString(ctx, raw, "function_version");
  PyDict_CopyItemString(ctx, raw, "invoked_function_arn");
  PyDict_CopyItemString(ctx, raw, "memory_limit_in_mb");
  PyDict_CopyItemString(ctx, raw, "aws_request_id");
  PyDict_CopyItemString(ctx, raw, "log_group_name");
  PyDict_CopyItemString(ctx, raw, "log_stream_name");

  PyObject* raw_identity = PyObject_GetAttrString(raw, "identity");
  PyObject* identity = PyDict_New();
  if (raw_identity != Py_None) {
    PyDict_CopyItemString(identity, raw_identity, "cognito_identity_id");
    PyDict_CopyItemString(identity, raw_identity, "cognito_identity_pool_id");
  }
  PyDict_SetSomeItemString(ctx, "identity", identity);
  Py_DECREF(raw_identity);

  PyObject* raw_client_context = PyObject_GetAttrString(raw, "client_context");
  PyObject* client_context = PyDict_New();
  if (raw_client_context != Py_None) {
    PyObject* raw_client = PyObject_GetAttrString(raw_client_context, "client");
    PyObject* client = PyDict_New();
    if (raw_client != Py_None) {
      PyDict_CopyItemString(client, raw_client, "installation_id");
      PyDict_CopyItemString(client, raw_client, "app_title");
      PyDict_CopyItemString(client, raw_client, "app_version_name");
      PyDict_CopyItemString(client, raw_client, "app_version_code");
      PyDict_CopyItemString(client, raw_client, "app_package_name");
    }
    PyDict_SetSomeItemString(client_context, "client", client);
    Py_DECREF(raw_client);

    PyDict_CopyItemString(client_context, raw_client_context, "env");
    PyDict_CopyItemString(client_context, raw_client_context, "custom");
  }
  PyDict_SetSomeItemString(ctx, "client_context", client_context);
  Py_DECREF(raw_client_context);

  PyObject* str = PyObject_CallMethod(py_json_module, "dumps", "O", ctx);

  PyDict_Free(ctx);

  return str;
}

static PyObject*
proxy_marshal_env()
{
  PyObject* environ = PyObject_GetAttrString(py_os_module, "environ");
  PyObject* dict = PyObject_GetAttrString(environ, "__dict__");
  PyObject* env = PyDict_GetItemString(dict, "data");

  PyObject* str = PyObject_CallMethod(py_json_module, "dumps", "O", env);

  Py_DECREF(environ);
  Py_DECREF(dict);

  return str;
}

static PyObject*
proxy_handle(PyObject* self, PyObject* args)
{
  PyObject *west = NULL;
  PyObject *evt = NULL;
  PyObject *ctx = NULL;

  PyArg_ParseTuple(args, "OO", &evt, &ctx);

  proxy_log_fn = PyObject_GetAttrString(ctx, "log");
  proxy_get_remaining_time_in_millis_fn = PyObject_GetAttrString(ctx, "get_remaining_time_in_millis");

  PyObject* jevt = PyObject_CallMethod(py_json_module, "dumps", "O", evt);
  PyObject* jctx = proxy_marshal_ctx(ctx);
  PyObject* jenv = proxy_marshal_env();

  char* sevt = PyString_AsString(jevt);
  char* sctx = PyString_AsString(jctx);
  char* senv = PyString_AsString(jenv);

  struct handle_return east = handle(sevt, sctx, senv);

  if (east.r0 != NULL) {
    PyObject* tmp = PyString_FromString(east.r0);
    west = PyObject_CallMethod(py_json_module, "loads", "O", tmp);
    Py_DECREF(tmp);
    free(east.r0);
  } else if (east.r1 != NULL) {
    PyErr_SetString(module_error, east.r1);
    free(east.r1);
    west = NULL;
  } else {
    Py_INCREF(Py_None);
    west = Py_None;
  }

  Py_DECREF(proxy_log_fn);
  Py_DECREF(proxy_get_remaining_time_in_millis_fn);
  Py_DECREF(jevt);
  Py_DECREF(jctx);
  Py_DECREF(jenv);

  return west;
}

#define STR(s) #s
#define XSTR(s) STR(s)
#define INIT_MODULE_SIG(name) PyMODINIT_FUNC init ## name (void)
#define INIT_MODULE(name) INIT_MODULE_SIG(name)

#ifndef FUNCTION
#define FUNCTION handler
#endif

#ifndef HANDLER
#define HANDLER handle
#endif

static PyMethodDef
module_methods[] = {
  {XSTR(HANDLER), (PyCFunction)proxy_handle, METH_VARARGS},
  {NULL, NULL}
};

INIT_MODULE(FUNCTION)
{
  py_json_module = PyImport_ImportModule("json");
  py_os_module = PyImport_ImportModule("os");
  PyObject *module = Py_InitModule(XSTR(FUNCTION), module_methods);
  module_error = PyErr_NewException(XSTR(FUNCTION)".error", NULL, NULL);
  Py_INCREF(module_error);
  PyModule_AddObject(module, "error", module_error);
}

#ifdef __cplusplus
}
#endif
