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