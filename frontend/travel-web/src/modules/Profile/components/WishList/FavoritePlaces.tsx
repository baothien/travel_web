import { Col, Pagination, PaginationProps, Row } from "antd";
import React, { useEffect, useState } from "react";
import { FavoriteItem } from "./FavoriteItem";
import { Place } from "../../../../models/place";
import placesApi from "../../../../apis/placesApi";

const FavoritePlaces: React.FC = () => {
  const [favoriteList, setFavoriteList] = useState<Place[]>([]);

  const [page, setPage] = React.useState<number>(1);

  const [totalRow, setTotalRow] = React.useState<number>(1);

  const onChangePaginate: PaginationProps["onChange"] = (page) => {
    setPage(page);
  };

  useEffect(() => {
    const params = {
      page: 1,
      limit: 10,
    };
    placesApi
      .getFavoriteList(params)
      .then((res) => {
        setFavoriteList(res.data.rows);
        setTotalRow(res.data.total_rows);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  return (
    <div>
      <Row gutter={[16, 16]} justify="space-between">
        {favoriteList.length === 0 && <h5>Chưa có địa điểm yêu thích</h5>}

        {favoriteList.map((item) => (
          <Col style={{ height: "100%" }} className="">
            <FavoriteItem
              key={item.id}
              id={item.id}
              thumbnail={item.thumbnail}
              name={item.name}
              placeType={item.place_type.name}
            />
          </Col>
        ))}
      </Row>

      <Pagination
        className="w-100 mt-4"
        onChange={onChangePaginate}
        total={totalRow}
        hideOnSinglePage={true}
      />
    </div>
  );
};

export default FavoritePlaces;
