#![allow(nonstandard_style)]
// Generated from Calc.g4 by ANTLR 4.8
use antlr_rust::tree::ParseTreeListener;
use super::calcparser::*;

pub trait CalcListener<'input> : ParseTreeListener<'input,CalcParserContextType>{

/**
 * Enter a parse tree produced by {@link CalcParser#prog}.
 * @param ctx the parse tree
 */
fn enter_prog(&mut self, _ctx: &ProgContext<'input>) { }
/**
 * Exit a parse tree produced by {@link CalcParser#prog}.
 * @param ctx the parse tree
 */
fn exit_prog(&mut self, _ctx: &ProgContext<'input>) { }

/**
 * Enter a parse tree produced by {@link CalcParser#expr}.
 * @param ctx the parse tree
 */
fn enter_expr(&mut self, _ctx: &ExprContext<'input>) { }
/**
 * Exit a parse tree produced by {@link CalcParser#expr}.
 * @param ctx the parse tree
 */
fn exit_expr(&mut self, _ctx: &ExprContext<'input>) { }

}
