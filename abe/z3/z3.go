// +build cgo,z3,!wasm

package z3

/*
#cgo CFLAGS: -std=gnu99
#cgo LDFLAGS: -lz3
#include <z3.h>
*/
import "C"

import (
  "runtime"
)

type Conf struct {
  p C.Z3_config
}

func freeConf(conf *Conf) {
  C.Z3_del_config(conf.p)
}

func NewConf() *Conf {
  conf := &Conf{p:C.Z3_mk_config()}
  runtime.SetFinalizer(conf, freeConf)
  return conf
}

type Ctx struct {
  p C.Z3_context
}

func freeCtx(ctx *Ctx) {
  C.Z3_del_context(ctx.p)
}

func NewCtx(conf *Conf) *Ctx {
  ctx := &Ctx{p:C.Z3_mk_context(conf.p)}
  runtime.SetFinalizer(ctx, freeCtx)
  return ctx
}

func (ctx *Ctx) EvalSmt2(s string) string {
  return C.GoString(C.Z3_eval_smtlib2_string(ctx.p, C.CString(s)))
}
