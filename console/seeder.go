package console

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/amacneil/dbmate/v2/pkg/driver/mysql"
	"github.com/notblessy/takeme-backend/db"
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var seederCmd = &cobra.Command{
	Use:   "seeder",
	Short: "seeder cmd",
	Long:  `This subcommand to execute seeding process`,
	Run:   seeder,
}

func init() {
	rootCmd.AddCommand(seederCmd)
}

func seeder(cmd *cobra.Command, args []string) {
	mysql := db.MysqlConnection()
	defer db.CloseMysql(mysql)

	subscriptionPlanRepo := repository.NewSubscriptionPlanRepository(mysql)

	err := seedSubscriptionPlan(subscriptionPlanRepo)
	continueOrFatal(err)

	log.Print("seed success")
}

func continueOrFatal(err error) {
	if err != nil {
		logrus.Fatal(err.Error())
	}
}

func seedSubscriptionPlan(planRepo model.SubscriptionPlanRepository) error {
	var subsPlan []model.SubscriptionPlan

	subs, err := os.Open("predefined/subscription_plan.json")
	if err != nil {
		return err
	}

	byteSubs, _ := ioutil.ReadAll(subs)

	err = json.Unmarshal(byteSubs, &subsPlan)
	if err != nil {
		return err
	}

	err = planRepo.BulkCreate(subsPlan)
	if err != nil {
		return err
	}

	return nil
}
