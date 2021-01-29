package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var createRepositoryCmd = &cobra.Command{
	Use:   "create::repository",
	Short: "Create repository commad",
	Long:  `This subcommand create Repository`,
	Run: func(cmd *cobra.Command, args []string) {

		// fmt.Println("create model" + args[0])

		var model string

		if len(args[0]) == 0 {
			fmt.Println("Please provide a valid command")
			return
		}

		input := input{
			param:       args[0],
			path:        "pkg",
			errorLine:   "model already exists! ",
			successLine: "File created successfully ",
		}

		model = input.param
		pathdirectory := input.path + "/" + strings.ToLower(input.param)

		// Check directory is exist
		stat := exists(pathdirectory)
		// If not exist create the directory
		if stat == false {
			// create the directory
			errx := os.Mkdir(input.path+"/"+strings.ToLower(input.param), 0755)
			if errx != nil {
				log.Fatal(errx)
			}

		}

		// TODO add basic model boilerplates
		service := pathdirectory + "/service.go"
		repository := pathdirectory + "/repository.go"

		//  var _, _ = os.Stat(filename)

		serviceExists := exists(service)

		var x string
		x = "package " + input.path + "\n\n" +
			"import (\n" +
			"\"github/sacsand/fiberPlus/model\"" + "\n" +
			"\"gorm.io/gorm\"" + "\n" +
			")\n\n" +

			"type Service interface { \n" +
			"   ReturnFinds(id int)(int,error) \n" +
			"   //other services... \n}" +
			"\n\n" +
			"//service struct \n" +
			"type service struct { \n" +
			"	repository Repository \n" +
			"} \n\n" +
			"// NewService is the single instance repo that is being created. \n" +
			"func NewService(r Repository) Service { \n" +
			"  return &service{ \n" +
			"    repository: r,\n" +
			"  }\n" +
			"}\n\n" +
			"//ReturnFinds - Return Finds\n" +
			"func (s *service) ReturnFinds(id int) (*model." + model + ", error) {\n" +
			"   return s.repository.ReturnFinds(id)\n" +
			"}"
		// If not exist create the directory
		if serviceExists == false {
			// create the directory
			createFile(service, x)
		}

		repositoryExists := exists(repository)

		if repositoryExists == false {

			var x string
			x = "package " + input.path + "\n\n" +
				"import (\n" +
				"\"github/sacsand/fiberPlus/model\"" + "\n" +
				"\"gorm.io/gorm\"" + "\n" +
				")\n\n" +

				"type Repository interface { \n" +
				"   Find(id int)(int,error) \n" +
				"   //other services... \n}" +
				"\n\n" +
				"//repository struct \n" +
				"type repository struct { \n" +
				"	db *gorm.DB \n" +
				"} \n\n" +
				"// NewRepo is the single instance repo that is being created. \n" +
				"func NewRepo(db *gorm.DB) Repository { \n" +
				"  return &repository{ \n" +
				"    db: db,\n" +
				"  }\n" +
				"}\n\n" +
				"//Find - Find user repository\n" +
				"func (r *repository) Find(id int) (*model." + model + ", error) {\n" +
				"   var data model." + model + "\n" +
				"   result := r.db.Where(\"id = ?\",id).Find(&data)\n" +
				"	if result.Error != nil {\n" +
				"      return data, result.Error\n" +
				"   }\n\n" +
				"   return data, nil\n" +
				"}"

			// create the directory
			createFile(repository, x)
		}

		fmt.Println(input.successLine, input.param)
		return

	},
}

func init() {
	RootCmd.AddCommand(createRepositoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
