import React, { useState, useEffect } from "react";
import checkInApi from "../../../../apis/checkInApi";

import "leaflet/dist/leaflet.css";

import "../../../Map/Map.scss";

import { MapContainer, TileLayer, ZoomControl, useMap } from "react-leaflet";

import { tileLayer } from "../../../../utils";

import L from "leaflet";

const center: [number, number] = [13.906325, 109.13269];

import ControllingGroup from "../../../Map/components/ControllingGroup";
import MapMarker from "./MapMarker";

const HistoryCheckIn: React.FC = () => {
  const [checkInItems, setCheckInItems] = useState<any>([]);

  useEffect(() => {
    const getData = async () => {
      await checkInApi.getCheckInList().then((res) => {
        const setPlacesId = new Set(res.data.rows.map((item) => item.place_id));

        const placesId = Array.from(setPlacesId);

        const dataCheckInPlace = placesId.map((placeId) => {
          const checkInImgUrls = res.data.rows
            .filter((item) => item.place_id === placeId)
            .map((item) => item.url);

          const placeInfo = res.data.rows.find(
            (item) => item.place_id === placeId
          ).place;

          return {
            place_id: placeId,
            checkInImgUrls,
            position: [placeInfo.lat, placeInfo.lng],
            title: placeInfo.name,
            thumbnail: placeInfo.thumbnail,
          };
        });

        setCheckInItems(dataCheckInPlace);
      });
    };

    getData();
  }, []);
  return (
    <div className="container-map">
      <MapContainer
        style={{ width: "46rem", height: "30rem" }}
        // className="vh-100 vw-100"
        zoomControl={false}
        center={center}
        zoom={10}
        scrollWheelZoom={true}
      >
        <GetCoordinates />

        {checkInItems.map((item) => (
          <MapMarker
            key={item.id}
            id={item.place_id}
            title={item.title}
            position={item.position}
            thumbnail={item.thumbnail}
            checkInImgUrls={item.checkInImgUrls}
          />
        ))}

        <TileLayer {...tileLayer} />

        <ControllingGroup />

        <ZoomControl position={"bottomright"} />
      </MapContainer>
    </div>
  );
};

export default HistoryCheckIn;

const GetCoordinates = () => {
  const map = useMap();

  useEffect(() => {
    if (!map) return;
    const info = L.DomUtil.create("div", "legend");

    const position = L.Control.extend({
      options: {
        position: "bottomleft",
      },

      onAdd: function () {
        info.textContent = "Click on map";
        return info;
      },
    });

    map.on("click", (e) => {
      info.textContent = e.latlng.toString();
    });

    map.addControl(new position());
  }, [map]);

  return null;
};
