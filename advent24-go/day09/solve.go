package day09

import (
	"slices"
	"strings"
)

type File struct {
	id, start, size int
}

type Free struct {
	start, size int
}

func parseInput(input string) ([]File, []Free) {
	files := make([]File, 0, (len(input)+1)/2)
	free := make([]Free, 0, len(input)/2)
	input = strings.TrimSpace(input)
	pos := 0
	fileId := 0
	for idx := 0; idx < len(input)-1; idx += 2 {
		size := int(input[idx] - '0')
		space := int(input[idx+1] - '0')
		files = append(files, File{fileId, pos, size})
		pos += size
		if space > 0 {
			free = append(free, Free{pos, space})
			pos += space
		}
		fileId++
	}
	size := int(input[len(input)-1] - '0')
	files = append(files, File{fileId, pos, size})
	return files, free
}

func getChecksum(id int, posStart int, size int) int {
	sum := 0
	for pos := posStart; pos < posStart+size; pos++ {
		sum += (id * pos)
	}
	return sum
}

func SolvePartOne(input string) int {
	files, frees := parseInput(input)
	checksum := 0

	fileIdx := len(files) - 1
	start := files[fileIdx].start
	fileSizeLeft := files[fileIdx].size

	freeIdx := 0
	pos := frees[freeIdx].start
	freeSizeLeft := frees[freeIdx].size

	for pos < start {
		if fileSizeLeft < freeSizeLeft {
			// It fits in the free space, space left
			checksum += getChecksum(files[fileIdx].id, pos, fileSizeLeft)

			pos += fileSizeLeft
			freeSizeLeft -= fileSizeLeft

			fileIdx--
			start = files[fileIdx].start
			fileSizeLeft = files[fileIdx].size

		} else if fileSizeLeft == freeSizeLeft {
			checksum += getChecksum(files[fileIdx].id, pos, fileSizeLeft)

			// It fits exactly
			fileIdx--
			start = files[fileIdx].start
			fileSizeLeft = files[fileIdx].size

			freeIdx++
			pos = frees[freeIdx].start
			freeSizeLeft = frees[freeIdx].size
		} else {
			// It does not fit in the free space
			checksum += getChecksum(files[fileIdx].id, pos, freeSizeLeft)

			fileSizeLeft -= freeSizeLeft

			freeIdx++
			pos = frees[freeIdx].start
			freeSizeLeft = frees[freeIdx].size
		}
	}

	if fileSizeLeft > 0 {
		checksum += getChecksum(files[fileIdx].id, files[fileIdx].start, fileSizeLeft)
	}

	for fidx := 0; fidx < fileIdx; fidx++ {
		checksum += getChecksum(files[fidx].id, files[fidx].start, files[fidx].size)
	}

	return checksum
}

func findIdx(frees []Free, fileSize int, filestart int) int {
	for j, f := range frees {
		if f.start >= filestart {
			break
		}
		if f.size >= fileSize {
			return j
		}
	}
	return -1
}

func SolvePartTwo(input string) int {
	files, frees := parseInput(input)
	checksum := 0
	for i := len(files) - 1; i >= 0; i-- {
		freeIdx := findIdx(frees, files[i].size, files[i].start)
		if freeIdx != -1 {
			checksum += getChecksum(files[i].id, frees[freeIdx].start, files[i].size)
			frees[freeIdx].start += files[i].size
			frees[freeIdx].size -= files[i].size
			if frees[freeIdx].size == 0 {
				frees = slices.Delete(frees, freeIdx, freeIdx+1)
			}
		} else {
			checksum += getChecksum(files[i].id, files[i].start, files[i].size)
		}
	}
	return checksum
}
