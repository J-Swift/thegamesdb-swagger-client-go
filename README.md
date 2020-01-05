# Go API client for gamesdb

API Documentations

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./gamesdb"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.thegamesdb.net*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DevelopersApi* | [**Developers**](docs/DevelopersApi.md#developers) | **Get** /v1/Developers | Fetch Developers list
*GamesApi* | [**GamesByGameID**](docs/GamesApi.md#gamesbygameid) | **Get** /v1/Games/ByGameID | Fetch game(s) by id
*GamesApi* | [**GamesByGameName**](docs/GamesApi.md#gamesbygamename) | **Get** /v1.1/Games/ByGameName | Fetch game(s) by name
*GamesApi* | [**GamesByGameNameV1**](docs/GamesApi.md#gamesbygamenamev1) | **Get** /v1/Games/ByGameName | Fetch game(s) by name
*GamesApi* | [**GamesByPlatformID**](docs/GamesApi.md#gamesbyplatformid) | **Get** /v1/Games/ByPlatformID | Fetch game(s) by platform id
*GamesApi* | [**GamesImages**](docs/GamesApi.md#gamesimages) | **Get** /v1/Games/Images | Fetch game(s) images by game(s) id
*GamesApi* | [**GamesUpdates**](docs/GamesApi.md#gamesupdates) | **Get** /v1/Games/Updates | Fetch games update
*GenresApi* | [**Genres**](docs/GenresApi.md#genres) | **Get** /v1/Genres | Fetch Genres list
*PlatformsApi* | [**Platforms**](docs/PlatformsApi.md#platforms) | **Get** /v1/Platforms | Fetch platforms list
*PlatformsApi* | [**PlatformsByPlatformID**](docs/PlatformsApi.md#platformsbyplatformid) | **Get** /v1/Platforms/ByPlatformID | Fetch platforms list by id
*PlatformsApi* | [**PlatformsByPlatformName**](docs/PlatformsApi.md#platformsbyplatformname) | **Get** /v1/Platforms/ByPlatformName | Fetch platforms by name
*PlatformsApi* | [**PlatformsImages**](docs/PlatformsApi.md#platformsimages) | **Get** /v1/Platforms/Images | Fetch platform(s) images by platform(s) id
*PublishersApi* | [**Publishers**](docs/PublishersApi.md#publishers) | **Get** /v1/Publishers | Fetch Publishers list


## Documentation For Models

 - [BaseApiResponse](docs/BaseApiResponse.md)
 - [Developer](docs/Developer.md)
 - [Developers](docs/Developers.md)
 - [DevelopersAllOf](docs/DevelopersAllOf.md)
 - [DevelopersAllOfData](docs/DevelopersAllOfData.md)
 - [Game](docs/Game.md)
 - [GameImage](docs/GameImage.md)
 - [GamesByGameId](docs/GamesByGameId.md)
 - [GamesByGameIdAllOf](docs/GamesByGameIdAllOf.md)
 - [GamesByGameIdAllOfData](docs/GamesByGameIdAllOfData.md)
 - [GamesByGameIdAllOfInclude](docs/GamesByGameIdAllOfInclude.md)
 - [GamesByGameIdAllOfIncludeBoxart](docs/GamesByGameIdAllOfIncludeBoxart.md)
 - [GamesByGameIdAllOfIncludePlatform](docs/GamesByGameIdAllOfIncludePlatform.md)
 - [GamesByGameIdV1](docs/GamesByGameIdV1.md)
 - [GamesByGameIdV1AllOf](docs/GamesByGameIdV1AllOf.md)
 - [GamesByGameIdV1AllOfInclude](docs/GamesByGameIdV1AllOfInclude.md)
 - [GamesImages](docs/GamesImages.md)
 - [GamesImagesAllOf](docs/GamesImagesAllOf.md)
 - [GamesImagesAllOfData](docs/GamesImagesAllOfData.md)
 - [GamesUpdates](docs/GamesUpdates.md)
 - [GamesUpdatesAllOf](docs/GamesUpdatesAllOf.md)
 - [GamesUpdatesAllOfData](docs/GamesUpdatesAllOfData.md)
 - [Genre](docs/Genre.md)
 - [Genres](docs/Genres.md)
 - [GenresAllOf](docs/GenresAllOf.md)
 - [GenresAllOfData](docs/GenresAllOfData.md)
 - [ImageBaseUrlMeta](docs/ImageBaseUrlMeta.md)
 - [PaginatedApiResponse](docs/PaginatedApiResponse.md)
 - [PaginatedApiResponseAllOf](docs/PaginatedApiResponseAllOf.md)
 - [PaginatedApiResponseAllOfPages](docs/PaginatedApiResponseAllOfPages.md)
 - [Platform](docs/Platform.md)
 - [PlatformImage](docs/PlatformImage.md)
 - [PlatformSkinny](docs/PlatformSkinny.md)
 - [Platforms](docs/Platforms.md)
 - [PlatformsAllOf](docs/PlatformsAllOf.md)
 - [PlatformsAllOfData](docs/PlatformsAllOfData.md)
 - [PlatformsByPlatformId](docs/PlatformsByPlatformId.md)
 - [PlatformsByPlatformName](docs/PlatformsByPlatformName.md)
 - [PlatformsByPlatformNameAllOf](docs/PlatformsByPlatformNameAllOf.md)
 - [PlatformsByPlatformNameAllOfData](docs/PlatformsByPlatformNameAllOfData.md)
 - [PlatformsImages](docs/PlatformsImages.md)
 - [PlatformsImagesAllOf](docs/PlatformsImagesAllOf.md)
 - [PlatformsImagesAllOfData](docs/PlatformsImagesAllOfData.md)
 - [Publisher](docs/Publisher.md)
 - [Publishers](docs/Publishers.md)
 - [PublishersAllOf](docs/PublishersAllOf.md)
 - [PublishersAllOfData](docs/PublishersAllOfData.md)
 - [UpdateModel](docs/UpdateModel.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



