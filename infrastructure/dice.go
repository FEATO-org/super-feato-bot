package infrastructure

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/FEATO-org/support-feato-system/domain/model"
	"github.com/FEATO-org/support-feato-system/domain/repository"
)

type DiceRepository struct {
}

func NewDiceRepository() repository.DiceRepository {
	return &DiceRepository{}
}

func (dr *DiceRepository) Roll(dice *model.Dice) (*model.Dice, error) {
	array, err := diceRoll(dice.GetQuery())
	if err != nil {
		return nil, err
	}
	total := int(sumArray(array))
	var formula string
	for index, a := range array {
		formula += strconv.Itoa(int(a))
		if (index + 1) != len(array) {
			formula += " + "
		}
	}
	dice.Set(dice.GetQuery(), formula, total)
	return dice, nil
}

// ダイスを振り、結果を配列で返す
func diceRoll(query string) ([]int64, error) {
	// すべて加算の前提のためダイスや足す数の区切りの判別に使う
	diceArray := strings.Split(query, "+")
	rand.Seed(time.Now().UnixNano())
	// 各ダイスの結果を格納する
	var calcArray []int64

	for _, val := range diceArray {
		// ダイスであるか、足す数かどうか判別する
		// ダイスであればダイスを振り次のループに行く
		if strings.Contains(val, "d") {
			roll := strings.Split(val, "d")
			if len(roll) > 2 {
				return nil, errors.New("Error!　dの数が多いです")
			}
			dice, err2 := strconv.Atoi(roll[1])
			count, err1 := strconv.Atoi(roll[0])
			if err1 != nil || err2 != nil {
				return nil, errors.New("Error!　数字以外のものが指定されました")
			}
			if count == 0 || dice == 0 {
				return nil, errors.New("Error!　ダイスの個数は1以上を指定してください")
			}
			for i := 0; i < count; i++ {
				calcArray = append(calcArray, rand.Int63n(int64(dice))+1)
			}
			continue
		}
		// ダイスでなければcalcArrayに数字を追加する
		sum, err := strconv.Atoi(val)
		if err != nil {
			return nil, errors.New("Error!　数字以外のものが指定されました")
		}
		calcArray = append(calcArray, int64(sum))
	}
	return calcArray, nil
}

// 与えられた配列の数字の合計を返す
func sumArray(array []int64) int64 {
	var result int64
	for _, a := range array {
		result += a
	}
	return result
}
