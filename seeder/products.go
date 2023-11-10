package seeder

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

var Products = []*model.Products{
	{
		Name:              "Minyak Goreng Refill Rose Brand 2L",
		Description:       "Minyak Goreng Rose Brand terbuat dari kelapa sawit pilihan berkualitas, diproses secara modern dengan teknologi tinggi secara higienis untuk membuat semya masakan menjadi lebih gurih dan lezat. Minyak Goreng Rose Brand mengandung BETA Karoten, omega 9, vitamin A dan E yang baik untuk tubuh.",
		CategoryID:        5,
		HazardousMaterial: false,
		Weight:            decimal.NewFromInt(22),
		Size:              decimal.NewFromInt(30),
		IsNew:             true,
		InternalSKU:       "OAKO OEasEF",
		ViewCount:         0,
		IsActive:          true,
		SellerID:          1,
	},
	{
		Name:              "Schneider Electric Leona Saklar Lampu - 2 Gang 2 Arah - LNA0600321",
		Description:       "Desain stylish dan minimalis untuk semua desain rumah Leona memiliki karakter berbentuk melingkar di setiap ujungnya. Desain yang tak lekang waktu dan sesuai untuk segala jenis rumah, serta memiliki berbagai varian untuk berbagai jenis kebutuhan, mulai dari saklar lampu, stop kontak schuko, tv, telepon dan data, hingga peredup lampu (dimmer). 2 Cara Pemasangan, sistem pencakar atau sekrup Saklar lampu dan stop kontak Leona hadir dengan 2 pilihan cara pemasangan. Dengan sistem pencakar dan sekrup yang memungkinkan untuk proyek renovasi maupun rumah baru. Sistem pencakar yang terlindungi, menjamin kekuatan dan daya cengkeram pada inbowdoost. Harga yang terjangkau dan kualitas terbaik Saklar lampu dan stop kontak Leona terbuat dari bahan polycarbonate berkualitas dan lebih aman karena lebih tahan panas, serta diperkuat dengan frame modul dari bahan logam yang menjamin kualitas, kekuatan dan tahan lebih lama.",
		CategoryID:        2,
		HazardousMaterial: false,
		Weight:            decimal.NewFromInt(22),
		Size:              decimal.NewFromInt(30),
		IsNew:             true,
		InternalSKU:       "OAKO OEKFOEF",
		ViewCount:         0,
		IsActive:          true,
		SellerID:          1,
	},
	{
		Name:              "Magsafe 2 Charger macbook 45w l 60w AIR l PRO - 45W",
		Description:       "MAGSAFE 2 LAGI PROMO MINGGU INI SILAHKAN ORDER SEBELUM HARGA KEMBALI NORMAL!!",
		CategoryID:        3,
		HazardousMaterial: false,
		Weight:            decimal.NewFromInt(22),
		Size:              decimal.NewFromInt(30),
		IsNew:             true,
		InternalSKU:       "OAKO",
		ViewCount:         0,
		IsActive:          true,
		SellerID:          3,
	},
}
