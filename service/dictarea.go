package service

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc/metadata"
	"gorm.io/gorm"

	"github.com/rocwg/dict-area-provider/model"
	pb "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

type DictAreaServiceServer struct {
	pb.UnimplementedDictAreaServiceServer
	DB *gorm.DB
}

// 1. 批量获取指定的字典数据
func (s *DictAreaServiceServer) GetDictByTypes(ctx context.Context, req *pb.GetDictByTypesRequest) (*pb.GetDictByTypesResponse, error) {
	if len(req.GetTypeCodes()) == 0 {
		return &pb.GetDictByTypesResponse{}, nil
	}

	var dbList []model.DictData
	// 查询启用的字典
	err := s.DB.Where("type_code IN ? AND is_enabled = ?", req.GetTypeCodes(), true).
		Order("sort_order ASC").Find(&dbList).Error
	if err != nil {
		return nil, err
	}

	// 组装返回的 Map
	dictMap := make(map[string]*pb.DictTypeResult)
	for _, item := range dbList {
		if _, ok := dictMap[item.TypeCode]; !ok {
			dictMap[item.TypeCode] = &pb.DictTypeResult{TypeCode: item.TypeCode}
		}
		dictMap[item.TypeCode].List = append(dictMap[item.TypeCode].List, &pb.DictData{
			DictLabel: item.DictLabel,
			DictValue: item.DictValue,
			SortOrder: item.SortOrder,
			CssClass:  item.CssClass,
		})
	}

	return &pb.GetDictByTypesResponse{DictMap: dictMap}, nil
}

// 2. 根据父级编码获取子级列表 (懒加载)
func (s *DictAreaServiceServer) GetAreaByParent(ctx context.Context, req *pb.GetAreaByParentRequest) (*pb.GetAreaByParentResponse, error) {
	parentCode := req.GetParentCode()
	if parentCode == "" {
		parentCode = "0" // 默认查省级
	}

	var dbList []model.Area
	err := s.DB.Where("parent_code = ? AND is_enabled = ?", parentCode, true).Find(&dbList).Error
	if err != nil {
		return nil, err
	}

	var list []*pb.AreaNode
	for _, item := range dbList {
		list = append(list, &pb.AreaNode{
			AreaCode:   item.AreaCode,
			ParentCode: item.ParentCode,
			AreaName:   item.AreaName,
			AreaLevel:  item.AreaLevel,
			MergerName: item.MergerName,
		})
	}

	//
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf(
			"metadata user_id=%v tenant_id=%v request_id=%v trace_id=%v",
			md.Get("x-user-id"),
			md.Get("x-tenant-id"),
			md.Get("x-request-id"),
			md.Get("x-trace-id"),
		)
	}

	return &pb.GetAreaByParentResponse{List: list}, nil
}

// 3. 根据编码批量反查地址信息
func (s *DictAreaServiceServer) BatchGetAreaByCodes(ctx context.Context, req *pb.BatchGetAreaByCodesRequest) (*pb.BatchGetAreaByCodesResponse, error) {
	if len(req.GetAreaCodes()) == 0 {
		return &pb.BatchGetAreaByCodesResponse{}, nil
	}

	var dbList []model.Area
	err := s.DB.Where("area_code IN ?", req.GetAreaCodes()).Find(&dbList).Error
	if err != nil {
		return nil, err
	}

	areaMap := make(map[string]*pb.AreaNode)
	for _, item := range dbList {
		areaMap[item.AreaCode] = &pb.AreaNode{
			AreaCode:   item.AreaCode,
			ParentCode: item.ParentCode,
			AreaName:   item.AreaName,
			AreaLevel:  item.AreaLevel,
			MergerName: item.MergerName,
		}
	}

	return &pb.BatchGetAreaByCodesResponse{AreaMap: areaMap}, nil
}

// 4. 模糊查询地址（简单实现版，后面有空再优于拼音）
func (s *DictAreaServiceServer) SearchArea(ctx context.Context, req *pb.SearchAreaRequest) (*pb.SearchAreaResponse, error) {
	keyword := req.GetKeyword()
	if keyword == "" {
		return &pb.SearchAreaResponse{}, nil
	}

	limit := int(req.GetLimit())
	if limit <= 0 || limit > 50 {
		limit = 20 // 默认限流 20 条防刷
	}

	var dbList []model.Area
	// 简单使用 LIKE 模糊查询名称
	err := s.DB.Where("area_name LIKE ? AND is_enabled = ?", "%"+keyword+"%", true).
		Limit(limit).Find(&dbList).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	var list []*pb.AreaNode
	for _, item := range dbList {
		list = append(list, &pb.AreaNode{
			AreaCode:   item.AreaCode,
			ParentCode: item.ParentCode,
			AreaName:   item.AreaName,
			AreaLevel:  item.AreaLevel,
			MergerName: item.MergerName,
		})
	}

	//
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf(
			"metadata user_id=%v tenant_id=%v request_id=%v trace_id=%v",
			md.Get("x-user-id"),
			md.Get("x-tenant-id"),
			md.Get("x-request-id"),
			md.Get("x-trace-id"),
		)
	}

	return &pb.SearchAreaResponse{List: list}, nil
}
