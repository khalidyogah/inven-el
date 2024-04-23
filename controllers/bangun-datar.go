package controllers

import (
	"net/http"
	. "q3/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HitungBangunDatar(ctx *gin.Context) {
	bangun := ctx.Param("bangun")

	switch {
	case bangun == "segitiga-sama-sisi":
		HitungSegitiga(ctx)
	case bangun == "persegi":
		HitungPersegi(ctx)
	case bangun == "persegi-panjang ":
		HitungPersegiPanjang(ctx)
	case bangun == "lingkaran ":
		HitungLingkaran(ctx)
	default:
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Perintah hitung tidak ada",
		})
	}
}
func HitungSegitiga(ctx *gin.Context) {
	var data Segitiga
	var out string

	data.Alas, _ = strconv.Atoi(ctx.Query("alas"))
	data.Tinggi, _ = strconv.Atoi(ctx.Query("tinggi"))
	mode := ctx.Query("hitung")

	if mode == "keliling" {
		out = strconv.FormatFloat(data.Keliling(), 'f', -1, 64)
	} else if mode == "luas" {
		out = strconv.FormatFloat(data.Luas(), 'f', -1, 64)

	} else {
		out = "Perintah hitung tidak ada"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": out,
	})
}
func HitungPersegi(ctx *gin.Context) {
	var data Persegi
	var out string

	data.Sisi, _ = strconv.Atoi(ctx.Query("sisi"))
	mode := ctx.Query("hitung")

	if mode == "keliling" {
		out = strconv.FormatFloat(data.Keliling(), 'f', -1, 64)
	} else if mode == "luas" {
		out = strconv.FormatFloat(data.Luas(), 'f', -1, 64)

	} else {
		out = "Perintah hitung tidak ada"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": out,
	})
}
func HitungPersegiPanjang(ctx *gin.Context) {
	var data PersegiPanjang
	var out string

	data.Panjang, _ = strconv.Atoi(ctx.Query("panjang"))
	data.Panjang, _ = strconv.Atoi(ctx.Query("lebar"))
	mode := ctx.Query("hitung")

	if mode == "keliling" {
		out = strconv.FormatFloat(data.Keliling(), 'f', -1, 64)
	} else if mode == "luas" {
		out = strconv.FormatFloat(data.Luas(), 'f', -1, 64)

	} else {
		out = "Perintah hitung tidak ada"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": out,
	})
}
func HitungLingkaran(ctx *gin.Context) {
	var data Lingkaran
	var out string

	data.JariJari, _ = strconv.Atoi(ctx.Query("jariJari"))
	mode := ctx.Query("hitung")

	if mode == "keliling" {
		out = strconv.FormatFloat(data.Keliling(), 'f', -1, 64)
	} else if mode == "luas" {
		out = strconv.FormatFloat(data.Luas(), 'f', -1, 64)

	} else {
		out = "Perintah hitung tidak ada"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": out,
	})
}
