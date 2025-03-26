package products

import (
	mock_products "codebase-service/mock/repository/products"
	model "codebase-service/models"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestProductsService(t *testing.T) {
	suite.Run(t, new(ProductServiceTestSuite))
}

type ProductServiceTestSuite struct {
	suite.Suite
	productRepo *mock_products.MockProductRepo
	service     ProductSvc
}

func (s *ProductServiceTestSuite) SetupTest() {
	s.productRepo = mock_products.NewMockProductRepo()
	s.service = NewProductSvc(s.productRepo)
}

func (s *ProductServiceTestSuite) TestGetProducts_Success() {
	req := new(model.GetProductsReq)
	res := new(model.GetProductsResp)

	s.productRepo.On("GetProducts", req).Return(res, nil)

	resp, err := s.service.GetProducts(req)

	s.NoError(err)
	s.NotNil(resp)
	s.productRepo.AssertExpectations(s.T())
}

func (s *ProductServiceTestSuite) TestGetProducts_Failed() {
	req := new(model.GetProductsReq)

	s.productRepo.On("GetProducts", req).Return(nil, sql.ErrConnDone)

	resp, err := s.service.GetProducts(req)

	s.Error(err)
	s.Nil(resp)
	s.productRepo.AssertExpectations(s.T())
}

func (s *ProductServiceTestSuite) TestGetProduct_Success() {
	req := new(model.GetProductReq)
	res := new(model.GetProductResp)

	s.productRepo.On("GetProduct", req).Return(res, nil)

	resp, err := s.service.GetProduct(req)

	s.NoError(err)
	s.NotNil(resp)
	s.productRepo.AssertExpectations(s.T())
}

func (s *ProductServiceTestSuite) TestGetProduct_NotFound() {
	req := new(model.GetProductReq)

	s.productRepo.On("GetProduct", req).Return(nil, sql.ErrNoRows)

	resp, err := s.service.GetProduct(req)

	s.Equal(sql.ErrNoRows, err)
	s.Nil(resp)
	s.productRepo.AssertExpectations(s.T())
}

func (s *ProductServiceTestSuite) TestGetProduct_Failed() {
	req := new(model.GetProductReq)

	s.productRepo.On("GetProduct", req).Return(nil, sql.ErrConnDone)

	resp, err := s.service.GetProduct(req)

	s.Error(err)
	s.Nil(resp)
	s.productRepo.AssertExpectations(s.T())
}

func (s *ProductServiceTestSuite) TestDeleteProduct_Success() {
	req := new(model.DeleteProductReq)

	s.productRepo.On("DeleteProduct", req).Return(nil)

	err := s.service.DeleteProduct(req)

	s.NoError(err)
	s.productRepo.AssertExpectations(s.T())
}

func (s *ProductServiceTestSuite) TestDeleteProduct_Failed() {
	req := new(model.DeleteProductReq)

	s.productRepo.On("DeleteProduct", req).Return(sql.ErrConnDone)

	err := s.service.DeleteProduct(req)

	s.Error(err)
	s.productRepo.AssertExpectations(s.T())
}

func (s *ProductServiceTestSuite) TestCreateProduct_Success() {
	req := new(model.CreateProductReq)
	res := new(model.GetProductResp)

	s.productRepo.On("CreateProduct", req).Return(res, nil)
	s.productRepo.On("IsShopOwner", req.UserId, req.ShopId).Return(nil)

	resp, err := s.service.CreateProduct(req)

	s.NoError(err)
	s.NotNil(resp)
	s.productRepo.AssertExpectations(s.T())
}

func (s *ProductServiceTestSuite) TestCreateProduct_Failed() {
	req := new(model.CreateProductReq)

	s.productRepo.On("CreateProduct", req).Return(nil, sql.ErrConnDone)
	s.productRepo.On("IsShopOwner", req.UserId, req.ShopId).Return(nil)

	resp, err := s.service.CreateProduct(req)

	s.Error(err)
	s.Nil(resp)
	s.productRepo.AssertExpectations(s.T())
}

func (s *ProductServiceTestSuite) TestCreateProduct_NotFound() {
	req := new(model.CreateProductReq)

	s.productRepo.On("CreateProduct", req).Return(nil, sql.ErrNoRows)
	s.productRepo.On("IsShopOwner", req.UserId, req.ShopId).Return(nil)

	resp, err := s.service.CreateProduct(req)

	s.Error(err)
	s.Nil(resp)
	s.productRepo.AssertExpectations(s.T())
}
