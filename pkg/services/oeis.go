package services

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/kasperaldrin/go-oeis/pkg/models"
)

func ParseOEIS(r io.Reader) (*models.OEISSequence, error) {
	seq := &models.OEISSequence{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || !strings.HasPrefix(line, "%") {
			continue
		}

		// %X A000001 ...
		parts := strings.SplitN(line, " ", 3)
		if len(parts) < 2 {
			continue
		}

		tag := parts[0]
		id := parts[1]
		payload := ""
		if len(parts) == 3 {
			payload = strings.TrimSpace(parts[2])
		}

		if seq.ID == "" {
			seq.ID = id
		}

		switch tag {
		case "%I":
			fields := strings.Fields(payload)
			seq.AltIDs = append(seq.AltIDs, fields...)

		case "%S", "%T", "%U":
			nums := strings.Split(payload, ",")
			for _, n := range nums {
				n = strings.TrimSpace(n)
				if n == "" {
					continue
				}
				v, err := strconv.ParseInt(n, 10, 64)
				if err == nil {
					seq.Sequence = append(seq.Sequence, v)
				}
			}

		case "%N":
			seq.Name = payload

		case "%C":
			seq.Comments = append(seq.Comments, payload)

		case "%D":
			seq.References = append(seq.References, payload)

		case "%H":
			seq.Links = append(seq.Links, payload)

		case "%F":
			seq.Formulas = append(seq.Formulas, payload)

		case "%Y":
			seq.CrossRefs = append(seq.CrossRefs, payload)

		case "%E":
			seq.Extensions = append(seq.Extensions, payload)

		case "%e":
			seq.Examples = append(seq.Examples, payload)

		case "%K":
			seq.Keywords = strings.Split(payload, ",")

		case "%A":
			seq.Author = payload

		case "%O":
			fields := strings.Fields(payload)
			if len(fields) >= 2 {
				seq.OffsetA, _ = strconv.Atoi(fields[0])
				seq.OffsetB, _ = strconv.Atoi(fields[1])
			}

		case "%p", "%t", "%o":
			seq.Programs = append(seq.Programs, models.ProgramEntry{
				Language: map[string]string{
					"%p": "Maple",
					"%t": "Mathematica",
					"%o": "Other",
				}[tag],
				Code: payload,
			})
		}
	}

	return seq, scanner.Err()
}
