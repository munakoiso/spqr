// Code generated by goyacc -o yacc/spqrparser/sql.go -p yy yacc/spqrparser/sql.y. DO NOT EDIT.

//line yacc/spqrparser/sql.y:3

package spqrparser

import __yyfmt__ "fmt"

//line yacc/spqrparser/sql.y:4

import (
	"strconv"
)

//line yacc/spqrparser/sql.y:15
type yySymType struct {
	yys       int
	empty     struct{}
	statement Statement
	show      *Show
	kr        *KeyRange
	sh_col    *ShardingColumn
	kill      *Kill
	drop      *Drop
	lock      *Lock
	shutdown  *Shutdown
	unlock    *Unlock
	split     *SplitKeyRange
	move      *MoveKeyRange
	str       string
	byte      byte
	int       int
	bool      bool
}

const STRING = 57346
const COMMAND = 57347
const SHOW = 57348
const KILL = 57349
const POOLS = 57350
const STATS = 57351
const LISTS = 57352
const SERVERS = 57353
const CLIENTS = 57354
const DATABASES = 57355
const CREATE = 57356
const ADD = 57357
const DROP = 57358
const LOCK = 57359
const UNLOCK = 57360
const SHUTDOWN = 57361
const SPLIT = 57362
const MOVE = 57363
const SHARDING = 57364
const COLUMN = 57365
const KEY = 57366
const RANGE = 57367
const SHARDS = 57368
const KEY_RANGES = 57369
const BY = 57370
const FROM = 57371
const TO = 57372

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"STRING",
	"COMMAND",
	"SHOW",
	"KILL",
	"POOLS",
	"STATS",
	"LISTS",
	"SERVERS",
	"CLIENTS",
	"DATABASES",
	"CREATE",
	"ADD",
	"DROP",
	"LOCK",
	"UNLOCK",
	"SHUTDOWN",
	"SPLIT",
	"MOVE",
	"SHARDING",
	"COLUMN",
	"KEY",
	"RANGE",
	"SHARDS",
	"KEY_RANGES",
	"BY",
	"FROM",
	"TO",
	"';'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacc/spqrparser/sql.y:264

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 74

var yyAct = [...]int{
	56, 59, 68, 32, 37, 28, 35, 34, 33, 65,
	64, 71, 53, 52, 51, 18, 19, 50, 49, 48,
	46, 36, 38, 13, 23, 24, 25, 26, 20, 21,
	22, 45, 44, 60, 43, 42, 41, 47, 29, 31,
	57, 69, 55, 27, 1, 54, 12, 11, 10, 16,
	58, 6, 61, 62, 63, 17, 7, 15, 5, 40,
	14, 66, 4, 3, 9, 67, 8, 39, 30, 70,
	2, 72, 0, 73,
}

var yyPact = [...]int{
	9, -1000, -26, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 16, -1000, -1000, -1000, -1000, -5, -5,
	-1000, 12, 11, 10, 8, 7, -4, -1000, -1000, 14,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -6, -7, -8, -11, -12, -13, 38, 36, 36,
	29, 36, 36, 36, -1000, -1000, -19, -1000, -21, 29,
	-1000, -1000, -1000, -1000, 36, 37, 37, -17, -1000, -1000,
	36, 29, -1000, -1000,
}

var yyPgo = [...]int{
	0, 70, 68, 67, 66, 64, 63, 62, 60, 58,
	57, 56, 55, 51, 49, 48, 47, 46, 39, 45,
	2, 1, 0, 44, 43,
}

var yyR1 = [...]int{
	0, 23, 24, 24, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 18, 18, 18, 18, 18, 18,
	18, 2, 3, 4, 19, 21, 6, 22, 20, 9,
	13, 7, 11, 8, 10, 14, 12, 16, 5, 17,
	15,
}

var yyR2 = [...]int{
	0, 2, 0, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 1, 1, 4, 1, 1, 1,
	1, 1, 1, 7, 4, 4, 4, 8, 2, 6,
	1,
}

var yyChk = [...]int{
	-1000, -23, -1, -6, -7, -9, -13, -11, -4, -5,
	-15, -16, -17, 14, -8, -10, -14, -12, 6, 7,
	19, 20, 21, 15, 16, 17, 18, -24, 31, 22,
	-2, -18, 8, 13, 12, 11, 26, 9, 27, -3,
	-18, 24, 24, 24, 24, 24, 24, 23, 25, 25,
	25, 25, 25, 25, -19, 4, -22, 4, -22, -21,
	4, -22, -22, -22, 29, 30, -21, -22, -20, 4,
	-20, 28, -22, -21,
}

var yyDef = [...]int{
	0, -2, 2, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 0, 31, 29, 30, 32, 0, 0,
	40, 0, 0, 0, 0, 0, 0, 1, 3, 0,
	23, 21, 14, 15, 16, 17, 18, 19, 20, 38,
	22, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 26, 24, 0, 27, 0, 0,
	25, 34, 35, 36, 0, 0, 0, 0, 39, 28,
	0, 0, 33, 37,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 31,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line yacc/spqrparser/sql.y:86
		{
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:87
		{
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:92
		{
			setParseTree(yylex, yyDollar[1].sh_col)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:96
		{
			setParseTree(yylex, yyDollar[1].kr)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:100
		{
			setParseTree(yylex, yyDollar[1].drop)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:104
		{
			setParseTree(yylex, yyDollar[1].lock)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:108
		{
			setParseTree(yylex, yyDollar[1].unlock)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:112
		{
			setParseTree(yylex, yyDollar[1].show)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:116
		{
			setParseTree(yylex, yyDollar[1].kill)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:120
		{
			setParseTree(yylex, yyDollar[1].shutdown)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:124
		{
			setParseTree(yylex, yyDollar[1].split)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:128
		{
			setParseTree(yylex, yyDollar[1].move)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:143
		{
			switch v := string(yyDollar[1].str); v {
			case ShowDatabasesStr, ShowPoolsStr, ShowShardsStr, ShowKeyRangesStr:
				yyVAL.str = v
			default:
				yyVAL.str = ShowUnsupportedStr
			}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:154
		{
			switch v := string(yyDollar[1].str); v {
			case KillClientsStr:
				yyVAL.str = v
			default:
				yyVAL.str = "unsupp"
			}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/spqrparser/sql.y:166
		{
			yyVAL.show = &Show{Cmd: yyDollar[2].str}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:173
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:179
		{
			yyVAL.int, _ = strconv.Atoi(string(yyDollar[1].str))
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/spqrparser/sql.y:185
		{
			yyVAL.sh_col = &ShardingColumn{ColName: yyDollar[4].str}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:191
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:198
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 33:
		yyDollar = yyS[yypt-7 : yypt+1]
//line yacc/spqrparser/sql.y:217
		{
			yyVAL.kr = &KeyRange{From: yyDollar[4].int, To: yyDollar[5].int, ShardID: yyDollar[6].str, KeyRangeID: yyDollar[7].str}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/spqrparser/sql.y:223
		{
			yyVAL.drop = &Drop{KeyRangeID: yyDollar[4].str}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/spqrparser/sql.y:229
		{
			yyVAL.lock = &Lock{KeyRangeID: yyDollar[4].str}
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/spqrparser/sql.y:235
		{
			yyVAL.unlock = &Unlock{KeyRangeID: yyDollar[4].str}
		}
	case 37:
		yyDollar = yyS[yypt-8 : yypt+1]
//line yacc/spqrparser/sql.y:242
		{
			yyVAL.split = &SplitKeyRange{KeyRangeID: yyDollar[4].str, KeyRangeFromID: yyDollar[6].str, Border: yyDollar[8].int}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/spqrparser/sql.y:248
		{
			yyVAL.kill = &Kill{Cmd: yyDollar[2].str}
		}
	case 39:
		yyDollar = yyS[yypt-6 : yypt+1]
//line yacc/spqrparser/sql.y:254
		{
			yyVAL.move = &MoveKeyRange{KeyRangeID: yyDollar[4].str, DestShardID: yyDollar[5].str}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/spqrparser/sql.y:260
		{
			yyVAL.shutdown = &Shutdown{}
		}
	}
	goto yystack /* stack new state and value */
}
