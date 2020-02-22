#script for generate language service pb
protoc proto/language/language-service.proto --go_out=plugins=grpc:.

#script for generate country service pb
protoc proto/country/country-service.proto --go_out=plugins=grpc:.

#script for generate category service pb
protoc proto/category/category-service.proto --go_out=plugins=grpc:.

#script for generate actor service pb
protoc proto/actor/actor-service.proto --go_out=plugins=grpc:.

#script for generate actor service pb
protoc proto/film_text/film_text-service.proto --go_out=plugins=grpc:.

#script for generate inventory service pb
protoc proto/inventory/inventory-service.proto --go_out=plugins=grpc:.

#script for generate film actor service pb
protoc proto/film/film-service.proto --go_out=plugins=grpc:.

#script for generate film category service pb
protoc proto/film_actor/film_actor-service.proto --go_out=plugins=grpc:.

#script for generate film category service pb
protoc proto/film_category/film-category-service.proto --go_out=plugins=grpc:.

#script for generate store service pb
protoc proto/store/store-service.proto --go_out=plugins=grpc:.

#script for generate staff service pb
protoc proto/staff/staff-service.proto --go_out=plugins=grpc:.

#script for generate customer service pb
protoc proto/customer/customer-service.proto --go_out=plugins=grpc:.

#script for generate payment service pb
protoc proto/payment/payment-service.proto --go_out=plugins=grpc:.

#script for generate rental service pb
protoc proto/rental/rental-service.proto --go_out=plugins=grpc:.

#script for generate city service pb
protoc proto/city/city-service.proto --go_out=plugins=grpc:.

#script for generate city service pb
protoc proto/address/address-service.proto --go_out=plugins=grpc:.