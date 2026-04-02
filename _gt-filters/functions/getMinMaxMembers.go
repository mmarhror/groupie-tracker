package functions

func GetMinMaxMembers(oneTree bool, fourSix bool, plusSeven bool) (int, int) {
	if oneTree && !fourSix && !plusSeven {
		return 1, 3
	}
	if !oneTree && fourSix && !plusSeven {
		return 4, 6
	}
	if !oneTree && !fourSix && plusSeven {
		return 7, 1000
	}
	if oneTree && fourSix && !plusSeven {
		return 1, 6
	}
	if oneTree && !fourSix && plusSeven {
		return 1, 1000
	}
	if !oneTree && fourSix && plusSeven {
		return 4, 1000
	}
	return 0, 0
}
