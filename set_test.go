/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package mapset

import (
	"testing"
)

func Test_NewSet(t *testing.T) {
	a := NewSet()

	if a.Size() != 0 {
		t.Error("NewSet should start out as an empty set")
	}
}

func Test_AddSet(t *testing.T) {
	a := NewSet()

	a.Add(1)
	a.Add(2)
	a.Add(3)

	if a.Size() != 3 {
		t.Error("AddSet does not have a size of 3 even though 3 items were added to a new set")
	}
}

func Test_AddSetNoDuplicate(t *testing.T) {
	a := NewSet()

	a.Add(7)
	a.Add(5)
	a.Add(3)
	a.Add(7)

	if a.Size() != 3 {
		t.Error("AddSetNoDuplicate set should have 3 elements since 7 is a duplicate")
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("AddSetNoDuplicate set should have a 7, 5, and 3 in it.")
	}
}

func Test_RemoveSet(t *testing.T) {
	a := NewSet()

	a.Add(6)
	a.Add(3)
	a.Add(1)

	a.Remove(3)

	if a.Size() != 2 {
		t.Error("RemoveSet should only have 2 items in the set")
	}

	if !(a.Contains(6) && a.Contains(1)) {
		t.Error("RemoveSet should have only items 6 and 1 in the set")
	}

	a.Remove(6)
	a.Remove(1)

	if a.Size() != 0 {
		t.Error("RemoveSet should be an empty set after removing 6 and 1")
	}
}

func Test_ContainsSet(t *testing.T) {
	a := NewSet()

	a.Add(71)

	if !a.Contains(71) {
		t.Error("ContainsSet should contain 71")
	}

	a.Remove(71)

	if a.Contains(71) {
		t.Error("ContainsSet should not contain 71")
	}

	a.Add(13)
	a.Add(7)
	a.Add(1)

	if !(a.Contains(13) && a.Contains(7) && a.Contains(1)) {
		t.Error("ContainsSet should contain 13, 7, 1")
	}

}

func Test_ClearSet(t *testing.T) {
	a := NewSet()

	a.Add(2)
	a.Add(5)
	a.Add(9)
	a.Add(10)

	a.Clear()

	if a.Size() != 0 {
		t.Error("ClearSet should be an empty set")
	}
}

func Test_SizeSet(t *testing.T) {
	a := NewSet()

	if a.Size() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(1)

	if a.Size() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Remove(1)

	if a.Size() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(9)

	if a.Size() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Clear()

	if a.Size() != 0 {
		t.Error("set should have a size of 1")
	}
}

func Test_SetIsSubset(t *testing.T) {
	a := NewSet()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(5)
	a.Add(7)

	b := NewSet()
	b.Add(3)
	b.Add(5)
	b.Add(7)

	if !b.IsSubset(a) {
		t.Error("set b should be a subset of set a")
	}

	b.Add(72)

	if b.IsSubset(a) {
		t.Error("set b should not be a subset of set a because it contains 72 which is not in the set of a")
	}

}

func Test_SetIsSuperSet(t *testing.T) {
	a := NewSet()
	a.Add(9)
	a.Add(5)
	a.Add(2)
	a.Add(1)
	a.Add(11)

	b := NewSet()
	b.Add(5)
	b.Add(2)
	b.Add(11)

	if !a.IsSuperset(b) {
		t.Error("set a should be a superset of set b")
	}

	b.Add(42)

	if a.IsSuperset(b) {
		t.Error("set a should not be a superset of set b because set a has a 42")
	}
}

func Test_SetUnion(t *testing.T) {
	a := NewSet()

	b := NewSet()
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	b.Add(5)

	c := a.Union(b)

	if c.Size() != 5 {
		t.Error("set c is unioned with an empty set and therefore should have 5 elements in it")
	}

	d := NewSet()
	d.Add(10)
	d.Add(14)
	d.Add(0)

	e := c.Union(d)
	if e.Size() != 8 {
		t.Error("set e should should have 8 elements in it after being unioned with set c to d")
	}

	f := NewSet()
	f.Add(14)
	f.Add(3)

	g := f.Union(e)
	if g.Size() != 8 {
		t.Error("set g should still ahve 8 elements in it after being unioned with set f that has duplicates")
	}
}

func Test_SetIntersect(t *testing.T) {
	a := NewSet()
	a.Add(1)
	a.Add(3)
	a.Add(5)

	b := NewSet()
	a.Add(2)
	a.Add(4)
	a.Add(6)

	c := a.Intersect(b)

	if c.Size() != 0 {
		t.Error("set c should be the empty set because there is no common items to intersect")
	}

	a.Add(10)
	b.Add(10)

	d := a.Intersect(b)

	if !(d.Size() == 1 && d.Contains(10)) {
		t.Error("set d should have a size of 1 and contain the item 10")
	}
}

func Test_SetDifference(t *testing.T) {
	a := NewSet()
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := NewSet()
	b.Add(1)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	b.Add(6)
	b.Add(99)

	c := a.Difference(b)

	if !(c.Size() == 1 && c.Contains(2)) {
		t.Error("the difference of set a to b is the set of 1 item: 2")
	}
}

func Test_SetSymmetricDifference(t *testing.T) {
	a := NewSet()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(45)

	b := NewSet()
	b.Add(1)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	b.Add(6)
	b.Add(99)

	c := a.SymmetricDifference(b)

	if !(c.Size() == 6 && c.Contains(2) && c.Contains(45) && c.Contains(4) && c.Contains(5) && c.Contains(6) && c.Contains(99)) {
		t.Error("the symmetric difference of set a to b is the set of 6 items: 2, 45, 4, 5, 6, 99")
	}
}

/*
func main() {

	a := NewSet()
	a.Add("Ralph")
	a.Add("John")
	a.Add("Caroline")
	a.Add("Paul")
	a.Add("Ralphie")

	b := NewSet()
	b.Add("Ralph")
	b.Add("Caroline")

	fmt.Println("IsSubset: ", b.IsSubset(a))
	fmt.Println("IsSuperset: ", a.IsSuperset(b))

	b.Add("Willis")
	fmt.Println("Union: ", a.Union(b))

	fmt.Println("Intersected: ", a.Intersect(b))

	fmt.Println("Differenced: ", a.Difference(b))

	fmt.Println("SymmetricDifferenced: ", a.SymmetricDifference(b))

	fmt.Println("Before Clear: ", a)
	a.Clear()
	fmt.Println("After Clear: ", a)

	//person test with set
	ralph := &Person{Name: "Ralph", Age: 34}
	ralphB := &Person{Name: "Ralph", Age: 34}
	sam := &Person{Name: "Sam", Age: 30}
	francois := &Person{Name: "Francois", Age: 38}
	peter := &Person{Name: "Peter", Age: 30}
	charles := &Person{Name: "Charles", Age: 40}
	chris := &Person{Name: "Chris", Age: 31}

	personSet := NewSet()
	personSet.Add(ralph)
	personSet.Add(sam)
	personSet.Add(francois)

	personSetB := NewSet()
	personSetB.Add(ralph)
	personSetB.Add(charles)
	personSetB.Add(francois)
	personSetB.Add(chris)
	personSetB.Add(peter)

	fmt.Println(personSet)
	fmt.Println(personSetB)

	fmt.Println(personSet.Difference(personSetB))

	//true comparing struct field-wise
	fmt.Println(*ralph == *ralphB)

	//false comparing pointers
	fmt.Println(ralph == ralphB)
}
*/

/*
type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf(`{"Name":"%s", "Age":%d}`, p.Name, p.Age)
}
*/
