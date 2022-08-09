package helpers

const DOMAIN = "192.168.1.81"

const PORT = "80"

const API_PREFIX = "/api"

const BACK_URL = "http://" + DOMAIN + ":" + PORT

const URL_BLOG_IMG = BACK_URL + API_PREFIX + "/blogs/%v/image"

const URL_PRODUCT_IMG = BACK_URL + API_PREFIX + "/products/image/%v?updatedAt=%v"
