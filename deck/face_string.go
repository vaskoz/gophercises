// Code generated by "stringer -type=Face"; DO NOT EDIT.

package deck

import "strconv"

const _Face_name = "AceTwoThreeFourFiveSixSevenEightNineTenJackQueenKing"

var _Face_index = [...]uint8{0, 3, 6, 11, 15, 19, 22, 27, 32, 36, 39, 43, 48, 52}

func (i Face) String() string {
	i -= 1
	if i >= Face(len(_Face_index)-1) {
		return "Face(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Face_name[_Face_index[i]:_Face_index[i+1]]
}