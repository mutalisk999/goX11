package goX11

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestCalcX11Hash(t *testing.T) {
	b1 := []byte("")
	b1r := CalcX11Hash(b1)
	fmt.Println("b1r: ", hex.EncodeToString(b1r))

	b2 := []byte("The quick brown fox jumps over the lazy dog")
	b2r := CalcX11Hash(b2)
	fmt.Println("b2r: ", hex.EncodeToString(b2r))

	b3 := []byte("DASH")
	b3r := CalcX11Hash(b3)
	fmt.Println("b3r: ", hex.EncodeToString(b3r))

	b4 := []byte("Take this kiss upon the brow! And, in parting from you now, Thus much let me avow-- You are not wrong, who deem That my days have been a dream; Yet if hope has flown away In a night, or in a day, In a vision, or in none, Is it therefore the less gone? All that we see or seem Is but a dream within a dream. I stand amid the roar Of a surf-tormented shore, And I hold within my hand Grains of the golden sand-- How few! yet how they creep Through my fingers to the deep, While I weep--while I weep! O God! can I not grasp Them with a tighter clasp? O God! can I not save One from the pitiless wave? Is all that we see or seem But a dream within a dream?")
	b4r := CalcX11Hash(b4)
	fmt.Println("b4r: ", hex.EncodeToString(b4r))
}

func TestCalcX11HashHex(t *testing.T) {
	b1 := []byte("")
	b1r := CalcX11HashHex(b1)
	fmt.Println("b1r: ", b1r)

	b2 := []byte("The quick brown fox jumps over the lazy dog")
	b2r := CalcX11HashHex(b2)
	fmt.Println("b2r: ", b2r)

	b3 := []byte("DASH")
	b3r := CalcX11HashHex(b3)
	fmt.Println("b3r: ", b3r)

	b4 := []byte("Take this kiss upon the brow! And, in parting from you now, Thus much let me avow-- You are not wrong, who deem That my days have been a dream; Yet if hope has flown away In a night, or in a day, In a vision, or in none, Is it therefore the less gone? All that we see or seem Is but a dream within a dream. I stand amid the roar Of a surf-tormented shore, And I hold within my hand Grains of the golden sand-- How few! yet how they creep Through my fingers to the deep, While I weep--while I weep! O God! can I not grasp Them with a tighter clasp? O God! can I not save One from the pitiless wave? Is all that we see or seem But a dream within a dream?")
	b4r := CalcX11HashHex(b4)
	fmt.Println("b4r: ", b4r)
}
