package target

import (
	"log"
	"testing"
	"time"
)

func simpleMh() MonthHemisphere {
	m := MonthHemisphere{
		Always: false,
		Ranges: "11-1,3-4,9",
		Array:  []int{11, 12, 1, 3, 4, 9}}

	m.WithText()
	return m
}

func Test_GenerateSequences(t *testing.T) {

	mn := MonthHemisphere{}
	// [1,2,3] => 0:[1,2,3]
	nums := []int{1, 2, 3}
	res := mn.generateSequences(nums)
	if len(res) != 1 || len(res[0]) != 3 {
		log.Fatalf("Failed to generate a simple sequence %v != %v\n", nums, res)
	}

	// [11, 12, 1, 2] => 0:[11,12,1,2]
	nums = []int{11, 12, 1, 2}
	res = mn.generateSequences(nums)
	if len(res) != 1 || len(res[0]) != 4 {
		log.Fatalf("Failed to generate a simple wrapping sequence %v != %v\n", nums, res)
	}

	// [1,2,5,6] => 0:[1,2], 1:[5,6]
	nums = []int{1, 2, 5, 6}
	res = mn.generateSequences(nums)
	if len(res) != 2 || len(res[0]) != 2 || len(res[1]) != 2 {
		log.Fatalf("Failed to generate a simple multi sequence %v != %v\n", nums, res)
	}

	// [11,12,1,5,6] => 0:[11,12,1], 1:[5,6]
	nums = []int{11, 12, 1, 5, 6}
	res = mn.generateSequences(nums)
	if len(res) != 2 || len(res[0]) != 3 || len(res[1]) != 2 {
		log.Fatalf("Failed to generate a wrapping multi sequence %v != %v\n", nums, res)
	}
	// [1,2, 11,12,1, 5] => 0[1,2], 1:[11,12,1], 2:[5]
	nums = []int{1, 2, 11, 12, 1, 5}
	res = mn.generateSequences(nums)
	if len(res) != 3 || len(res[0]) != 2 || len(res[1]) != 3 || len(res[2]) != 1 {
		log.Fatalf("Failed to generate a complex multi sequence %v != %v\n", nums, res)
	}

	// [12,1,2,3, 1, 6,7, 10,11,12] => 0:[12,1,2,3], 1:[1], 2:[6,7], 3:[10,11,12]
	nums = []int{12, 1, 2, 3, 1, 6, 7, 10, 11, 12}
	res = mn.generateSequences(nums)
	if len(res) != 4 || len(res[0]) != 4 || len(res[1]) != 1 || len(res[2]) != 2 || len(res[3]) != 3 {
		log.Fatalf("Failed to generate a complex multi sequence %v != %v\n", nums, res)
	}

}

func Test_SequenceToText(t *testing.T) {

	mn := MonthHemisphere{}
	// Feb
	nums := []int{2}
	seq := mn.generateSequences(nums)
	res := mn.sequencesToText(seq)
	if res != "February" {
		log.Fatalf("Failed to created the correct string for the sequence:\n%v\n%v\n", seq, res)
	}

	nums = []int{1, 3}
	seq = mn.generateSequences(nums)
	res = mn.sequencesToText(seq)
	if res != "January, March" {
		log.Fatalf("Failed to created the correct string for the sequence:\n%v\n%v\n", seq, res)
	}

	// [1,2,3] => 0:[1,2,3] => Jan -> Mar
	nums = []int{1, 2, 3}
	seq = mn.generateSequences(nums)
	res = mn.sequencesToText(seq)
	if res != "January - March" {
		log.Fatalf("Failed to created the correct string for the sequence:\n%v\n%v\n", seq, res)
	}

	nums = []int{1, 2, 4, 5}
	seq = mn.generateSequences(nums)
	res = mn.sequencesToText(seq)
	if res != "January - February, April - May" {
		log.Fatalf("Failed to created the correct string for the sequence:\n%v\n%v\n", seq, res)
	}

	nums = []int{11, 12, 1, 4, 5}
	seq = mn.generateSequences(nums)
	res = mn.sequencesToText(seq)
	if res != "November - January, April - May" {
		log.Fatalf("Failed to created the correct string for the sequence:\n%v\n%v\n", seq, res)
	}

	nums = []int{11, 12, 1, 4, 5, 11, 12}
	seq = mn.generateSequences(nums)
	res = mn.sequencesToText(seq)
	if res != "November - January, April - May, November - December" {
		log.Fatalf("Failed to created the correct string for the sequence:\n%v\n%v\n", seq, res)
	}

	nums = []int{11, 12, 1, 4, 5, 12}
	seq = mn.generateSequences(nums)
	res = mn.sequencesToText(seq)
	if res != "November - January, April - May, December" {
		log.Fatalf("Failed to created the correct string for the sequence:\n%v\n%v\n", seq, res)
	}

}

func Test_Is(t *testing.T) {

	jan := time.Date(2021, time.Month(1), 1, 1, 0, 0, 0, time.UTC)
	mh := simpleMh()
	is := mh.Is(jan)
	// based on simple type, not new, is leaving, is available
	if is.New || is.Leaving == false || is.Availabile == false {
		log.Fatalf("Failed to generate correct Is details:\n%v\n%v\n%v\n", jan.String(), mh.Array, is)
	}

	dec := time.Date(2021, time.Month(11), 1, 1, 0, 0, 0, time.UTC)
	is = mh.Is(dec)
	// based on simple type, is new, not leaving, is available
	if is.New == false || is.Leaving || is.Availabile == false {
		log.Fatalf("Failed to generate correct Is details:\n%v\n%v\n%v\n", dec.String(), mh.Array, is)
	}

	sept := time.Date(2021, time.Month(9), 1, 1, 0, 0, 0, time.UTC)
	is = mh.Is(sept)
	// based on simple type, is new, is leaving, is available
	if is.New == false || is.Leaving == false || is.Availabile == false {
		log.Fatalf("Failed to generate correct Is details:\n%v\n%v\n%v\n", sept.String(), mh.Array, is)
	}

	aug := time.Date(2021, time.Month(8), 1, 1, 0, 0, 0, time.UTC)
	is = mh.Is(aug)
	// based on simple type, not new, not leaving, not available
	if is.New || is.Leaving || is.Availabile {
		log.Fatalf("Failed to generate correct Is details:\n%v\n%v\n%v\n", aug.String(), mh.Array, is)
	}

}
