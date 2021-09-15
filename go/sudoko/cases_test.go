package main

type sudokuTest struct {
	input    string
	expected string
}

var testCases = []bracketTest{
	{
		"[]",
		true,
	},
	{
		"",
		true,
	},
	{
		"[[",
		false,
	},
	{
		"}{",
		false,
	},
	{
		"{]",
		false,
	},
	{
		"{ }",
		true,
	},
	{
		"{[])",
		false,
	},
	{
		"{[]}",
		true,
	},
	{
		"{}[]",
		true,
	},
	{
		"([{}({}[])])",
		true,
	},
	{
		"{[)][]}",
		false,
	},
	{
		"([{])",
		false,
	},
	{
		"[({]})",
		false,
	},
	{
		"{}[",
		false,
	},
	{
		"[]]",
		false,
	},
	{
		"(((185 + 223.85) * 15) - 543)/2",
		true,
	},
	{
		"\\left(\\begin{array}{cc} \\frac{1}{3} & x\\\\ \\mathrm{e}^{x} &... x^2 \\end{array}\\right)",
		true,
	},
}
