package verifier

import (
	"testing"

	"github.com/svenski123/POET/server/go/poet"
)

func TestVerifier(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	// debugLog.SetOutput(os.Stdout)
	// defer debugLog.SetOutput(ioutil.Discard)
	test_n := 4
	p := poet.NewProver(false) // Should declare Prover with n
	v := NewVerifier(p, test_n)
	b := []byte{'a', 'b'}
	err := v.Commit(b)
	if err != nil {
		t.Error("Error Sending Commitment: ", err)
	}
	_, err = v.GetCommitProof()
	if err != nil {
		t.Error("Error Getting Commit Proof: ", err)
	}
	_, err = v.SelectChallenge()
	if err != nil {
		t.Error("Error Selecting Challenge: ", err)
	}
	err = v.Challenge()
	if err != nil {
		t.Error("Error Getting Challenge: ", err)
	}
	_, err = v.GetChallengeProof()
	if err != nil {
		t.Error("Error Getting Challenge Proof: ", err)
	}
	//fmt.Println("Verifying Challenge Proof", v.challengeProof)
	err = v.VerifyChallengeProof()
	if err != nil {
		t.Error("Error Verifying Challenge Proof: ", err)
	}
}

func TestNIPVerifier(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	// debugLog.SetOutput(os.Stdout)
	// defer debugLog.SetOutput(ioutil.Discard)
	p := poet.NewProver(true)
	test_n := 4
	v := NewVerifier(p, test_n)
	b := []byte{'a', 'b'}
	err := v.Commit(b)
	if err != nil {
		t.Error("Error Sending Commitment: ", err)
	}
	_, err = v.GetCommitProof()
	if err != nil {
		t.Error("Error Getting Commit Proof: ", err)
	}
	_, err = v.GetChallengeProof()
	if err != nil {
		t.Error("Error Getting Challenge Proof: ", err)
	}
	//fmt.Println("Verifying Challenge Proof", v.challengeProof)
	err = v.VerifyChallengeProof()
	if err != nil {
		t.Error("Error Verifying Challenge Proof: ", err)
	}
}
