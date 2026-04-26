package easy

func furthestDistanceFromOrigin(moves string) int {
	cntL := 0
	cntR := 0
	anyW := 0
	for _, c := range moves {
		if c == 'L' {
			cntL += 1
		}else if c == 'R' {
			cntR += 1
		} else {
			anyW += 1
		}
	}
	if cntL > cntR {
		return cntL - cntR + anyW
	}else if cntR > cntL {
		return cntR - cntL + anyW
	} else {
		return anyW
	}
}