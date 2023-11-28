package util

import (
	"fmt"
)

type ErrorCode int

const (
	BadRequest ErrorCode = iota
	NotFound
	Unauthorized
)

type CustomError struct {
	Code    ErrorCode
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func New(code ErrorCode, message string) error {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrWrongCredentials                                         = &CustomError{Message: "wrong email or password", Code: Unauthorized}
	ErrNoRecordFound                                            = &CustomError{Message: "not found", Code: NotFound}
	ErrBadPIN                                                   = &CustomError{Message: "length minimum length is 6 and should be numeric", Code: BadRequest}
	ErrInvalidInput                                             = &CustomError{Message: "invalid input", Code: BadRequest}
	ErrWalletAlreadySet                                         = &CustomError{Message: "wallet already activated", Code: BadRequest}
	ErrWalletNotSet                                             = &CustomError{Message: "your wallet has not been activated", Code: BadRequest}
	ErrWalletPINNotMatch                                        = &CustomError{Message: "your wallet pin doesnt match", Code: BadRequest}
	ErrInvalidPassword                                          = &CustomError{Message: "invalid password", Code: Unauthorized}
	ErrSameWalletPIN                                            = &CustomError{Message: "your new wallet PIN must not be the same as your current PIN"}
	ErrEmailAlreadyExist                                        = &CustomError{Message: "email already registered", Code: BadRequest}
	ErrUsernameAlreadyExist                                     = &CustomError{Message: "username already registered", Code: BadRequest}
	ErrCantUseThisEmail                                         = &CustomError{Message: "you cannot use this email", Code: BadRequest}
	ErrPasswordContainUsername                                  = &CustomError{Message: "password cannot contains username as part of it", Code: BadRequest}
	ErrSameEmail                                                = &CustomError{Message: "you already used this email", Code: BadRequest}
	ErrInvalidAmountRange                                       = &CustomError{Message: "amount should be between 50000 and 10000000", Code: BadRequest}
	ErrSamePhoneNumber                                          = &CustomError{Message: "you already used this phone number", Code: BadRequest}
	ErrSameUsername                                             = &CustomError{Message: "you alrady use this username", Code: BadRequest}
	ErrCantUseThisUsername                                      = &CustomError{Message: "username already used", Code: BadRequest}
	ErrCantUseThisPhonenumber                                   = &CustomError{Message: "phone number already used", Code: BadRequest}
	ErrOrderStatusNotWaiting                                    = &CustomError{Message: "you can only cancel orders that are still awaiting seller confirmation", Code: BadRequest}
	ErrWeakPassword                                             = &CustomError{Message: "password should have at least one uppercase, one lowercase with minimum length is 8", Code: BadRequest}
	ErrOrderNotFound                                            = &CustomError{Message: "We could not find the order you are looking for or the order has been processed previously", Code: BadRequest}
	ErrInsufficientStock                                        = &CustomError{Message: "Insufficient product stock", Code: BadRequest}
	ErrInsufficientBalance                                      = &CustomError{Message: "Your wallet balance is insufficient, please top up first to proceed with order checkout", Code: BadRequest}
	ErrQtyInputZero                                             = &CustomError{Message: "Quantity must be minimum one", Code: BadRequest}
	ErrCourierNotAvailable                                      = &CustomError{Message: "courier not found or not available", Code: BadRequest}
	ErrInvalidDateFormat                                        = &CustomError{Message: "invalid date format. the date format is YYYY-MM-DD, e.g: 2023-09-10", Code: BadRequest}
	ErrWalletHistorySortBy                                      = &CustomError{Message: "sortBy should be one of 'amount', 'type', 'date'", Code: BadRequest}
	ErrProductFavoriteSortBy                                    = &CustomError{Message: "sortBy should be one of 'price', or 'date'", Code: BadRequest}
	ErrQtyExceed                                                = &CustomError{Message: "The quantity of product available currently is less than the quantity you requested", Code: BadRequest}
	ErrProductCartNotFound                                      = &CustomError{Message: "Product not found", Code: BadRequest}
	ErrAlreadyRegisteredAsSeller                                = &CustomError{Message: "already registered as seller", Code: BadRequest}
	ErrAddressNotAvailable                                      = &CustomError{Message: "address not found or available", Code: BadRequest}
	ErrInvalidToken                                             = &CustomError{Message: "Invalid token", Code: Unauthorized}
	ErrUnauthorize                                              = &CustomError{Message: "Unauthorize", Code: Unauthorized}
	ErrProductNotFound                                          = &CustomError{Message: "Product not found", Code: NotFound}
	ErrProductOrderNotFound                                     = &CustomError{Message: "The product order was not found or the order has not been completed", Code: BadRequest}
	ErrInvalidRating                                            = &CustomError{Message: "Ratings should be between 1 - 5", Code: BadRequest}
	ErrAlreadyReviewProduct                                     = &CustomError{Message: "You have given a review of this product", Code: BadRequest}
	EmailNotFound                                               = &CustomError{Message: "we can't find this email address", Code: BadRequest}
	ErrRequestForgetToken                                       = &CustomError{Message: "The link you used is invalid or unable to be processed", Code: Unauthorized}
	ErrSellerNotFound                                           = &CustomError{Message: "Seller not found", Code: NotFound}
	ErrOrderDetailNotFound                                      = &CustomError{Message: "We could not find the order you are looking for", Code: NotFound}
	ErrSamePassword                                             = &CustomError{Message: "New Password cannot be the same with old password", Code: BadRequest}
	ErrNoImage                                                  = &CustomError{Message: "You should choose at least one product image", Code: BadRequest}
	ErrIncorrectPassword                                        = &CustomError{Message: "Invalid Password", Code: Unauthorized}
	ErrOrderNotInDelivered                                      = &CustomError{Message: "The order you selected is not or has not been delivered", Code: BadRequest}
	ErrOrderNotCompleted                                        = &CustomError{Message: "The order you selected has not been completed", Code: BadRequest}
	ErrInvalidOTP                                               = &CustomError{Message: "Invalid OTP", Code: Unauthorized}
	ErrExpiredOTP                                               = &CustomError{Message: "OTP Expired, send a new OTP Request", Code: BadRequest}
	ErrPasswordIdentical                                        = &CustomError{Message: "New password identical to old password, please create a unique Password", Code: BadRequest}
	ErrShopPromotionNotFound                                    = &CustomError{Message: "shop promotion not found", Code: BadRequest}
	ErrMarketplacePromotionNotFound                             = &CustomError{Message: "marketplace promotion not found", Code: BadRequest}
	ErrShowcaseNotFound                                         = &CustomError{Message: "showcase not found", Code: BadRequest}
	ErrShopPromotionMinimumAmountGreaterThanTotalPayment        = &CustomError{Message: "shop promotion minimum amount is greater than total payment", Code: BadRequest}
	ErrMarketplacePromotionMinimumAmountGreaterThanTotalPayment = &CustomError{Message: "marketplace promotion minimum amount is greater than total payment", Code: BadRequest}
	ErrSameShop                                                 = &CustomError{Message: "you can't buy your own product", Code: BadRequest}
	ErrVariantPhotoFailed                                       = &CustomError{Message: "failed to upload variant photo. please try again and choose variant photo.", Code: BadRequest}
)
