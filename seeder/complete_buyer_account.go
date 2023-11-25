package seeder

import (
	"fmt"
	"math/rand"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func SeedBuyers(db *gorm.DB) {
	count := 0
	existingName := map[string]string{}

	for count < 20 {
		name, email := pickRandomName()

		if _, exist := existingName[name]; !exist {
			buyer := createBuyer(db, name, email)
			createBuyerAddress(db, buyer.ID)
			existingName[name] = name
			count++
		}
		fmt.Println("1")
	}
}

func createBuyer(db *gorm.DB, name string, email string) model.Accounts {
	numbers := map[string]string{}

	number := ""
	numberCreated := false
	for !numberCreated {
		n := generateRandomNumbers(1, 9, 11)

		if _, exist := numbers[n]; !exist {
			number = n
			numbers[n] = n
			numberCreated = true
		}
	}

	account := model.Accounts{

		FullName:       name,
		Username:       name,
		Email:          email + "@mail.com",
		PhoneNumber:    "+62" + number,
		Password:       "$2a$14$ggRGSX9uKrEfapylGVadWee/P1yCOKduFFqnzNdq7U3ble5nxtNqC",
		Gender:         randomGender(),
		Birthdate:      time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
		ProfilePicture: "https://mangathrill.com/wp-content/uploads/2019/07/Portgas.D..Ace_.full_.5794251280x720.png",
		WalletPin:      "123456",
		Balance:        decimal.NewFromInt(0),
	}

	if err := db.Create(&account).Error; err != nil {
		panic(err)
	}

	account.WalletNumber = fmt.Sprint(4200000000000 + account.ID)

	if err := db.Updates(&account).Error; err != nil {
		panic(err)
	}

	return account
}

func createBuyerAddress(db *gorm.DB, accountId int) {
	address := randomAdrress()

	address.AccountID = accountId
	address.IsSellerDefault = false

	if err := db.Create(&address).Error; err != nil {
		panic(err)
	}
}

func pickRandomName() (string, string) {
	ranFirstName := rand.Intn((len(firstName)-1)-0+1) + 0
	ranLastName := rand.Intn((len(lastName)-1)-0+1) + 0

	fn := firstName[ranFirstName]
	ln := lastName[ranLastName]

	return fn + " " + ln, fn + ln
}

var firstName = []string{"Ahmad", "Agus", "Adi", "Andi", "Arief", "Budi", "Bayu", "Bagus", "Candra", "Dedi", "Eko", "Fajar", "Gilang", "Hendra", "Heri", "Iwan", "Joko", "Kurniawan", "Luki", "Made", "Nanda", "Oka", "Putu", "Rizki", "Satria", "Teguh", "Ujang", "Vino", "Wahyu", "Yudi", "Zulkifli", "Aditya", "Arif", "Bambang", "Cecep", "Deni", "Edi", "Fahmi", "Guntur", "Hadi", "Indra", "Jamal", "Krisna", "Lukman", "Mahendra", "Nizar", "Opik", "Prasetyo", "Qori", "Rendy", "Syahputra", "Taufik", "Udin", "Wawan", "Yani", "Zainal", "Aldi", "Angga", "Bram", "Chandra", "Dika", "Eka", "Fadli", "Gunawan", "Habibi", "Ikbal", "Januar", "Kiki", "Leo", "Maman", "Nono", "Okta", "Pandu", "Radit", "Samsul", "Tedi", "Usman", "Vicky", "Wildan", "Yoga", "Zaki", "Alif", "Anton", "Benny", "Cakrawala", "Dody", "Elang", "Farhan", "Hendri", "Imam", "Jefri", "Kurnia", "Lutfi", "Munir", "Novi", "Ovi", "Purnama", "Rian", "Sandi", "Tio", "Anita", "Ayu", "Amalia", "Bunga", "Cinta", "Dewi", "Eka", "Farah", "Gita", "Hana", "Intan", "Jihan", "Kirana", "Lestari", "Maya", "Nadia", "Oka", "Putri", "Rina", "Sari", "Tika", "Umi", "Vita", "Wulan", "Yanti", "Zara", "Ani", "Bella", "Clara", "Dian", "Ella", "Fitri", "Gina", "Hesti", "Indah", "Junita", "Kania", "Lia", "Mira", "Nia", "Ola", "Puspita", "Rini", "Sinta", "Tia", "Ulfa", "Vina", "Winda", "Yulia", "Zahra", "Andini", "Bintang", "Citra", "Dina", "Elisa", "Fina", "Hilda", "Ika", "Juwita", "Kartika", "Lina", "Melati", "Nurul", "Oktavia", "Prita", "Risa", "Selvi", "Tiara", "Ulfah", "Wati", "Yuni", "Zulfa", "Anggun", "Budiarti", "Chika", "Desi", "Evi", "Fifi", "Hani", "Ida", "Jessy", "Karina", "Laila", "Mila", "Nina", "Okky", "Prima", "Retno", "Shinta", "Tri", "Uut", "Veni", "Windy", "Yola", "Zulfi", "Astri", "Bambang", "Cici", "Desta", "Eni"}
var lastName = []string{"Aditya", "Agustina", "Ahmad", "Aji", "Akbar", "Alamsyah", "Ali", "Amin", "Andrianto", "Anggraini", "Anjani", "Ardianto", "Arief", "Arifin", "Asmara", "Aswandi", "Atmaja", "Azhari", "Bachtar", "Bagaskara", "Bakti", "Bambang", "Bangun", "Baskoro", "Basri", "Bastian", "Batubara", "Bawono", "Bayu", "Budianto", "Budiman", "Cahyadi", "Cahyono", "Candra", "Darmadi", "Darmawan", "Darma", "Darmono", "Darsono", "Daryanto", "Daud", "Dewanto", "Eka", "Fadil", "Fadillah", "Fahmi", "Faisal", "Fajar", "Farid", "Fatimah", "Febrian", "Ferdi", "Galang", "Gani", "Gunadi", "Gunawan", "Gunarto", "Hadi", "Hafidz", "Halim", "Hamdan", "Handoko", "Harefa", "Harianto", "Haris", "Hartanto", "Hartono", "Hasan", "Hasim", "Hasnawi", "Hasyim", "Heri", "Hermanto", "Hermawan", "Heru", "Hidayat", "Hidayanto", "Huda", "Husaini", "Husni", "Idris", "Ilham", "Ilyas", "Indra", "Irawan", "Iskandar", "Ismail", "Iwan", "Irawan", "Jaelani", "Jamal", "Jatmiko", "Jaya", "Jayadi", "Kadir", "Kurnia", "Kusuma", "Laksana", "Laksono", "Latif", "Lesmana", "Lestari", "Ma'ruf", "Mahendra", "Mahesa", "Makmur", "Malik", "Mangunsong", "Mansur", "Margono", "Mariadi", "Martadinata", "Maryadi", "Mihardja", "Mulyadi", "Munir", "Mustafa", "Natalegawa", "Natsir", "Nawawi", "Nazar", "Nurdin", "Nurjaman", "Oktaviani", "Pambudi", "Pandu", "Pangestu", "Pardede", "Parto", "Perdana", "Prabowo", "Prakosa", "Pranata", "Prasetya", "Prasetyo", "Pratama", "Prayitno", "Priadi", "Priyanto", "Purnama", "Putra", "Putri", "Rahadi", "Rahardjo", "Rahmanto", "Rahman", "Rahmat", "Rais", "Raja", "Ramadhan", "Ramli", "Rasyid", "Ridwan", "Riswandi", "Rizal", "Rizki", "Rochman", "Rodiyanto", "Rojali", "Rosadi", "Rosman", "Rudiyanto", "Rusdi", "Ruslan", "Sadiq", "Sadono", "Sahar", "Said", "Salim", "Samad", "Samsudin", "Saputra", "Saputro", "Saragih", "Sari", "Satria", "Sastrawan", "Setiabudi", "Setiadi", "Setiawan", "Setyawan", "Sidik", "Sihombing", "Simbolon", "Sinaga", "Siregar", "Soebroto", "Soedjatmoko", "Soekarno", "Soepomo", "Soerjanto", "Soetanto", "Soetomo", "Somantri", "Suhaimi", "Sukarno", "Suleiman", "Sumantri", "Surya", "Susanto", "Susilo", "Sya'ban", "Syah", "Syamsuddin", "Syarif", "Tahir", "Tamba", "Tanjung", "Tanojo", "Tarihoran", "Tarmizi", "Taslim", "Thamrin", "Tirta", "Tjahjono", "Tjahyadi", "Trisna", "Usman", "Utama", "Utomo", "Wahid", "Wahyudi", "Wahyuni", "Wibawa", "Wibisono", "Wijaya", "Wijoyo", "Winata", "Winoto", "Wiryawan", "Wisnu", "Wiyono", "Yanto", "Yasin", "Yudha", "Yudhistira", "Yudhoyono", "Yulianto", "Yulius", "Yunus", "Yusuf", "Zaelani", "Zain", "Zakaria", "Zulkarnain", "Zulkifli", "Zulmi", "Zulqarnain", "Abdullah", "Abidin", "Aminah", "Aisyah", "Asri", "Astuti", "Aziz", "Bintoro", "Damayanti", "Darman", "Dewanti", "Fadhil", "Fadilah", "Farida", "Fatimah", "Halima", "Hasanah", "Hidayah", "Iman", "Iskandar", "Jalal", "Jamaludin", "Jamil", "Kamil", "Karim", "Khair", "Lathifah", "Latifah", "Lukman", "Ma'ruf", "Mahmud", "Malik", "Mansur", "Mardani", "Marwah", "Mas'ud", "Masykur", "Mawardi", "Naim", "Najib", "Nawal", "Nazaruddin", "Nazir", "Nuh", "Nurhayati", "Qadir", "Rahmani", "Rahmatullah", "Rais", "Rasyid", "Ridwan", "Saad"}
