package day300

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var (
	out io.Writer = os.Stdout // nolint
	in  io.Reader = os.Stdin  // nolint
)

// VotingResult is the number of votes a candidate received.
type VotingResult struct {
	Candidate, Votes int
}

func main() {
	var voterID, candidateID int
	voters := make(map[int]struct{})
	candidates := make(map[int]*VotingResult)
	top3 := make([]*VotingResult, 0, 3)
	_, err := fmt.Fscanf(in, "%d,%d", &voterID, &candidateID)
	for err != io.EOF {
		if _, voted := voters[voterID]; voted {
			fmt.Fprintf(out, "voter %d already voted. this vote for candidate %d will not count\n", voterID, candidateID)
		} else {
			voters[voterID] = struct{}{}
			if cand, exists := candidates[candidateID]; !exists {
				candidates[candidateID] = &VotingResult{Candidate: candidateID, Votes: 1}
				if len(top3) < 3 {
					top3 = append(top3, candidates[candidateID])
				}
			} else {
				cand.Votes++
				if cand.Votes > top3[0].Votes && !Contains(top3, candidates[candidateID]) {
					top3[0] = cand
				}
			}
			sort.Slice(top3, func(i, j int) bool {
				return top3[i].Votes < top3[j].Votes
			})
			for _, v := range top3 {
				fmt.Fprintf(out, "%d,%v ", v.Candidate, v.Votes)
			}
			fmt.Fprintln(out)
		}
		_, err = fmt.Fscanf(in, "%d,%d", &voterID, &candidateID)
	}
}

func Contains(top3 []*VotingResult, cand *VotingResult) bool {
	for _, c := range top3 {
		if cand == c {
			return true
		}
	}
	return false
}
