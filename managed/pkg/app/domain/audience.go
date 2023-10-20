package domain

import (
	"context"
	"github.com/viant/parsly"
	"github.com/viant/sqlparser"
	"github.com/viant/sqlparser/expr"
	"github.com/viant/sqlparser/node"
	"strconv"
	"strings"
)

type Deal struct {
	Id   int32   `sqlx:"name=ID"`
	Name *string `sqlx:"name=NAME"`
	Fee  *string `sqlx:"name=FEE"`
}

type Audience struct {
	Id              int32   `sqlx:"name=ID"`
	Name            *string `sqlx:"name=NAME"`
	MatchExpression *string `sqlx:"name=MATCH_EXPRESSION"`
	dealIds         []int   `sqlx:"-"`
	Deals           []*Deal
}

func (a *Audience) OnFetch(ctx context.Context) error {
	if a.MatchExpression != nil && *a.MatchExpression != "" {
		err := a.enrichDealIdsWithExpression()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Audience) enrichDealIdsWithExpression() error {
	qualify := expr.Qualify{}
	cursor := parsly.NewCursor("", []byte(*a.MatchExpression), 0)
	if err := sqlparser.ParseQualify(cursor, &qualify); err != nil {
		return err
	}

	sqlparser.Traverse(qualify.X, func(n node.Node) bool {
		switch actual := n.(type) {
		case *expr.Binary:
			x := sqlparser.Stringify(actual.X)
			if strings.ToLower(actual.Op) == "in" && strings.ToLower(x) == "deals" {
				par := actual.Y.(*expr.Parenthesis)
				values := par.Raw[1 : len(par.Raw)-1]
				for _, value := range strings.Split(values, ",") {
					value = strings.TrimSpace(value)
					if intVal, err := strconv.Atoi(value); err == nil {
						a.dealIds = append(a.dealIds, intVal)
					}
				}
				return true
			}
		}
		return true
	})
	return nil
}
