package gh

import (
	"bufio"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ckeyer/sloth/types"
	"gopkg.in/mgo.v2/bson"
)

type Project struct {
	types.Project `json:",inline" bson:",inline"`

	cli *Client
}

type CommitPatch types.CommitPatch

type FilePatch types.FilePatch

func (p *Project) GetCommitPatch(sha string) (*CommitPatch, []*FilePatch, error) {
	rc, resp, err := p.cli.Repositories.GetCommit(context.Background(), p.Owner, p.ReposName, sha)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, nil, fmt.Errorf("github server return %s", resp.Status)
	}

	cp := &CommitPatch{
		ProjectID: p.ID,
		SHA:       mustString(rc.SHA),
		Author:    mustString(rc.Author.Login),
		Message:   mustString(rc.Commit.Message),
	}
	fps := []*FilePatch{}

	for _, f := range rc.Files {
		oldL, newL, err := parseFilePatch(mustString(f.Patch))
		if err != nil {
			return nil, nil, err
		}

		fp := &FilePatch{
			ID:        bson.NewObjectId(),
			ProjectID: p.ID,
			CommitSHA: mustString(rc.SHA),
			Filename:  mustString(f.Filename),
			Old:       oldL,
			New:       newL,
		}
		fps = append(fps, fp)
	}

	return cp, fps, nil
}

// parseFilePatch parse "@@ -1,7 +1,7 @@"
func parseFilePatch(patch string) (oldL *types.LinePatch, newL *types.LinePatch, err error) {
	line, _, err := bufio.NewReader(strings.NewReader(patch)).ReadLine()
	if err != nil {
		return nil, nil, err
	}

	ftLine := string(line)
	if strings.HasPrefix(ftLine, "@@ ") && strings.HasSuffix(ftLine, " @@") {
		return parseLinePatch(ftLine)
	} else {
		return nil, nil, fmt.Errorf("not found @@ ... @@ at first line: %s", ftLine)
	}
}

func parseLinePatch(pairsLine string) (oldL *types.LinePatch, newL *types.LinePatch, err error) {
	pairsLine = strings.TrimPrefix(pairsLine, "@@ ")
	pairsLine = strings.TrimSuffix(pairsLine, " @@")

	changes := strings.Split(pairsLine, " ")
	if len(changes) != 2 {
		return nil, nil, fmt.Errorf("can not parse changes: %s", pairsLine)
	}

	// parse "-1,7"
	parsePair := func(pair string) (*types.LinePatch, error) {
		ps := strings.Split(pair, ",")
		if len(ps) != 2 {
			return nil, fmt.Errorf("invalid pair: %s", pair)
		}

		offset, err := strconv.Atoi(ps[0])
		if err != nil {
			return nil, err
		}
		count, err := strconv.Atoi(ps[1])
		if err != nil {
			return nil, err
		}

		return &types.LinePatch{
			Offset: offset,
			Count:  count,
		}, nil
	}

	oldL, err = parsePair(changes[0])
	if err != nil {
		return nil, nil, err
	}
	newL, err = parsePair(changes[1])
	if err != nil {
		return nil, nil, err
	}

	return
}
