package service

import (
	"errors"

	"github.com/diskordanz/consumer/service/model"
)

func (s ConsumerService) GetFranchise(id int) (model.Franchise, error) {
	franchise, err := s.fh.GetFranchise(id)
	if err != nil {
		return model.Franchise{}, err
	}
	resultFranchise := model.MapToFranchise(franchise)
	return resultFranchise, nil
}

func (s ConsumerService) ListFranchises(count, offset int) ([]model.Franchise, error) {
	list, err := s.fh.ListFranchises(count, offset)
	if err != nil {
		return nil, err
	}
	result := model.MapToFranchiseList(list)
	return result, nil
}

/////////////////////////////////////////////////////

func (s ConsumerService) GetLocationsOfFranchise(franchiseID, count, offset int) ([]model.Location, error) {
	locations, err := s.lh.GetLocationsOfFranchise(franchiseID, count, offset)
	if err != nil {
		return nil, err
	}
	result := model.MapToLocationsList(locations)
	return result, nil
}

/*
func (s ConsumerService) GetLocationsOfFranchiseByName(franchiseID, count, offset int, name string) ([]model.Location, error) {
	locations, err := s.lh.GetLocationsOfFranchiseByName(franchiseID, count, offset, name)
	if err != nil {
		return nil, err
	}
	result := model.MapToLocationsList(locations)
	return result, nil
}*/

/////////////////////////////////////////////////////

func (s ConsumerService) ListCategories() ([]model.Category, error) {
	list, err := s.ch.ListCategories()
	if err != nil {
		return nil, err
	}
	result := model.MapToCategoryList(list)
	return result, nil
}

/////////////////////////////////////////////////////

func (s ConsumerService) ListOrders(id, count, offset int) ([]model.Order, error) {
	list, err := s.oh.ListOrders(id, count, offset)
	if err != nil {
		return nil, err
	}
	result := model.MapToOrderList(list)
	return result, nil
}

func (s ConsumerService) GetOrder(id int) (model.Order, error) {
	order, err := s.oh.GetOrder(id)
	if err != nil {
		return model.Order{}, err
	}
	resultOrder := model.MapToOrder(order)
	return resultOrder, nil
}

/////////////////////////////////////////////////////////////

func (s ConsumerService) GetProduct(id int) (model.Product, error) {
	product, err := s.ph.GetProductById(id)
	if err != nil {
		return model.Product{}, err
	}
	resultProduct := model.MapToProduct(product)
	return resultProduct, nil
}

func (s ConsumerService) ListProducts(name string, count, offset int) ([]model.Product, error) {
	products, err := s.ph.ListProducts(name, count, offset)
	if err != nil {
		return nil, err
	}
	result := model.MapToProductList(products)
	return result, nil
}

func (s ConsumerService) ListProductsByCategory(id uint64, name string, count, offset int) ([]model.Product, error) {
	products, err := s.ph.ListProductsByCategory(id, name, count, offset)
	if err != nil {
		return nil, err
	}
	result := model.MapToProductList(products)
	return result, nil
}

////////////////////////////////////////////////////////////////

func (s ConsumerService) GetConsumer(id int) (model.Consumer, error) {
	consumer, err := s.uh.GetConsumer(id)
	if err != nil {
		return model.Consumer{}, err
	}
	resultConsumer := model.MapToConsumer(consumer)
	return resultConsumer, nil
}

func (s ConsumerService) CreateConsumer(consumer model.Consumer) (model.Consumer, error) {
	dbConsumer := model.MapToConsumerDB(consumer)
	c, err := s.uh.CreateConsumer(dbConsumer)
	if err != nil {
		return model.Consumer{}, err
	}
	return model.MapToConsumer(c), nil
}

func (s ConsumerService) UpdateConsumer(consumer model.Consumer) (model.Consumer, error) {
	dbConsumer := model.MapToConsumerDB(consumer)
	dbConsumer, err := s.uh.UpdateConsumer(dbConsumer)
	if err != nil {
		return model.Consumer{}, err
	}
	return model.MapToConsumer(dbConsumer), nil
}

func (s ConsumerService) Healthcheck() error {
	if s.fh == nil || s.uh == nil || s.lh == nil || s.ch == nil || s.ph == nil || s.oh == nil {
		return errors.New("error: problem with database handlers")
	}
	return nil
}
