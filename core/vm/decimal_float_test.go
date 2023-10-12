// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"fmt"
	"testing"

	"github.com/holiman/uint256"
)

// func TestInt(t *testing.T) {

// 	tests := []struct {
// 		a uint256.Int
// 		b uint256.Int
// 	}{
// 		{*uint256.NewInt(20), *uint256.NewInt(7)},
// 		{*uint256.NewInt(20), *new(uint256.Int).Neg(uint256.NewInt(7))},
// 		{*new(uint256.Int).Neg(uint256.NewInt(20)), *uint256.NewInt(7)},
// 		{*new(uint256.Int).Neg(uint256.NewInt(20)), *new(uint256.Int).Neg(uint256.NewInt(7))},
// 	}
// 	for _, tt := range tests {

// 		fmt.Println("a", showInt(&tt.a))
// 		fmt.Println("b", showInt(&tt.b))

// 		var out uint256.Int

// 		out.Add(&tt.a, &tt.b)
// 		fmt.Println("Add", showInt(&out))

// 		out.Sub(&tt.a, &tt.b)
// 		fmt.Println("Sub", showInt(&out))

// 		out.Div(&tt.a, &tt.b)
// 		fmt.Println("Div", showInt(&out))

// 		signed_div(&tt.a, &tt.b, &out)
// 		fmt.Println("signed_div", showInt(&out))

// 		out.Mul(&tt.a, &tt.b)
// 		fmt.Println("Mul", showInt(&out))

// 		out.Exp(&tt.a, &tt.b)
// 		fmt.Println("Exp", showInt(&out))

// 		out.Mod(&tt.a, &tt.b)
// 		fmt.Println("Mod", showInt(&out))

// 		out.Abs(&tt.a)
// 		fmt.Println("Abs", showInt(&out))

// 		out.Neg(&tt.a)
// 		fmt.Println("Neg", showInt(&out))
// 	}
// }

// func TestAdd(t *testing.T) {
// 	tests := []struct {
// 		a decimal
// 		b decimal
// 		c decimal
// 	}{
// 		{decimal{*uint256.NewInt(5), *ZERO_uint256_Int}, decimal{*uint256.NewInt(121), *MINUS_ONE_uint256_Int}, decimal{*uint256.NewInt(171), *MINUS_ONE_uint256_Int}},
// 		{decimal{*uint256.NewInt(5), *ZERO_uint256_Int}, decimal{*uint256.NewInt(121), *ZERO_uint256_Int}, decimal{*uint256.NewInt(126), *ZERO_uint256_Int}},
// 		{decimal{*new(uint256.Int).Neg(uint256.NewInt(2)), *MINUS_ONE_uint256_Int}, decimal{*uint256.NewInt(8), *MINUS_ONE_uint256_Int}, decimal{*uint256.NewInt(6), *MINUS_ONE_uint256_Int}},
// 	}
// 	for _, tt := range tests {
// 		var out decimal
// 		add(&tt.a, &tt.b, &out, false)
// 		fmt.Println("a", showDecimal(&tt.a), "b", showDecimal(&tt.b), "out", showDecimal(&out), "c", showDecimal(&tt.c))
				
// 		var out2 decimal
// 		normalize(&out, &out2, 0, true, false)

// 		if out2 != tt.c {
// 			t.Fatal(tt.a, tt.b, out, out2, tt.c)
// 		}
// 	}
// }

// func TestMultiply(t *testing.T) {
// 	tests := []struct {
// 		a decimal
// 		b decimal
// 		c decimal
// 	}{
// 		{decimal{*uint256.NewInt(2), *ZERO_uint256_Int}, decimal{*uint256.NewInt(2), *ZERO_uint256_Int}, decimal{*uint256.NewInt(4), *ZERO_uint256_Int}},
// 		{decimal{*uint256.NewInt(2), *ZERO_uint256_Int}, decimal{*uint256.NewInt(5), MINUS_ONE_uint256_Int}, decimal{*uint256.NewInt(1), *ZERO_uint256_Int}},
// 	}
// 	for _, tt := range tests {
// 		var out decimal
// 		multiply(&tt.a, &tt.b, &out, false)
// 		normalize(&out, &out, 10, false, false)
// 		fmt.Println("a", tt.a, "b", tt.b, "out", out, "c", tt.c)
// 		// if out != tt.c {
// 		// 	t.Fatal(tt.a, tt.b, out, tt.c)
// 		// }
// 	}
// }

// func TestDiv(t *testing.T) {
// 	tests := []struct {
// 		a decimal
// 		b decimal
// 		c decimal
// 	}{
// 		{decimal{*ONE_uint256_Int, *TEN_uint256_Int}, decimal{*ONE_uint256_Int, *ZERO_uint256_Int}, decimal{*ONE_uint256_Int, *TEN_uint256_Int}},
// 	}
// 	for _, tt := range tests {
// 		var out decimal
// 		divide(&tt.a, &tt.b, &out, true)
// 		fmt.Println("a", tt.a, "b", tt.b, "out", out, "c", tt.c)
// 		// if out != tt.c {
// 		// 	t.Fatal(tt.a, tt.b, out, tt.c)
// 		// }
// 	}
// }

// func TestNormalize(t *testing.T) {

// 	LARGE_TEN := uint256.NewInt(10)
// 	LARGE_TEN.Exp(LARGE_TEN, uint256.NewInt(76))
// 	NEG_77 := new(uint256.Int).Neg(uint256.NewInt(76))

// 	tests := []struct {
// 		a decimal
// 		b decimal
// 	}{
// 		// {decimal{*ONE_uint256_Int, *ZERO_uint256_Int}, decimal{*ONE_uint256_Int, *ZERO_uint256_Int}},
// 		{decimal{*uint256.NewInt(100), *new(uint256.Int).Neg(uint256.NewInt(2))}, ONE},
// 		{decimal{*LARGE_TEN, *NEG_77}, ONE},
// 	}
// 	for _, tt := range tests {
// 		var out decimal

// 		normalize(&tt.a, &out, 100, false, false)
// 		showDecimal("a", &tt.a)
// 		showDecimal("out", &out)
// 		showDecimal("b", &tt.b)

// 		if out != tt.b {
// 			t.Fatal(tt.a, tt.b, out)
// 		}
// 	}
// }

// func TestInv(t *testing.T) {

// 	tests := []struct {
// 		a decimal
// 		b decimal
// 	}{
// 		{decimal{*ONE_uint256_Int, *ZERO_uint256_Int}, decimal{*ONE_uint256_Int, *ZERO_uint256_Int}},
// 		{decimal{*uint256.NewInt(2), *ZERO_uint256_Int}, decimal{*uint256.NewInt(5), MINUS_ONE_uint256_Int}},
// 		{decimal{*new(uint256.Int).Neg(uint256.NewInt(20)), MINUS_ONE_uint256_Int}, decimal{*new(uint256.Int).Neg(uint256.NewInt(5)), MINUS_ONE_uint256_Int}},
// 	}
// 	for _, tt := range tests {

// 		showDecimal("a", &tt.a)

// 		var out decimal

// 		inverse(&tt.a, &out, false)
// 		// showDecimal("a", &tt.a)
// 		// showDecimal("out", &out)
// 		// showDecimal("b", &tt.b)

// 		var out2 decimal
// 		multiply(&out, &tt.a, &out2, false)
// 		// showDecimal("out2", &out2)

// 		var out3 decimal
// 		normalize(&out2, &out3, 0, false, false)
// 		// showDecimal("out3", &out3)

// 		if out3 != ONE {
// 			t.Fatal(tt.a, out, tt.b)
// 		}
// 	}
// }

// func TestExp(t *testing.T) {
// 	tests := []struct {
// 		a decimal
// 		b decimal
// 	}{
// 		{decimal{*ONE_uint256_Int, *ZERO_uint256_Int}, decimal{*uint256.NewInt(2718281), *new(uint256.Int).Neg(uint256.NewInt(6))}},
// 	}
// 	for _, tt := range tests {
// 		var out decimal
// 		steps := uint(10)
// 		exp(&tt.a, &out, steps, false)
// 		showDecimal("a", &tt.a)
// 		showDecimal("out", &out)
// 		showDecimal("b", &tt.b)
// 		// if out != tt.c {
// 		// 	t.Fatal(tt.a, tt.b, out, tt.c)
// 		// }
// 	}
// }

// func TestSubtract(t *testing.T) {
// 	tests := []struct {
// 		a decimal
// 		b decimal
// 		c decimal
// 	}{
// 		{decimal{*uint256.NewInt(2), *ZERO_uint256_Int},decimal{*uint256.NewInt(5), *MINUS_ONE_uint256_Int}, decimal{*uint256.NewInt(15), *MINUS_ONE_uint256_Int}},
// 	}
// 	for _, tt := range tests {
		
// 		var out decimal
// 		subtract(&tt.a, &tt.b, &out, true)
// 		fmt.Println("a", showDecimal(&tt.a))
// 		fmt.Println("b", showDecimal(&tt.b))
// 		fmt.Println("out", showDecimal(&out))

// 		var out2 decimal
// 		normalize(&out, &out2, 0, true, false)
// 		fmt.Println("out2", showDecimal(&out2))

// 		if out2 != tt.c {
// 			t.Fatal(tt.a, tt.b, out, out2, tt.c)
// 		}
// 	}
// }

func TestLt(t *testing.T) {
	tests := []struct {
		a decimal
		b decimal
		c bool
	}{
		{decimal{*uint256.NewInt(5), *ZERO_uint256_Int}, decimal{*uint256.NewInt(2), *ZERO_uint256_Int}, false},
		{decimal{*uint256.NewInt(5), *MINUS_ONE_uint256_Int}, decimal{*uint256.NewInt(2), *ZERO_uint256_Int}, true},
	}
	for _, tt := range tests {
		fmt.Println("a", showDecimal(&tt.a))
		fmt.Println("b", showDecimal(&tt.b))
		lt := lessthan(&tt.a, &tt.b, true)
		if lt != tt.c {
			t.Fatal(tt.a, tt.b, tt.c)
		}
	}
}
