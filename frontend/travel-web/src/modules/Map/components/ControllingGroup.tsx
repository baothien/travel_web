import React from "react";
import { latLng, LatLng, LocationEvent } from "leaflet";
import L from "leaflet";
import { LayersControl, useMapEvents } from "react-leaflet";

const center: [number, number] = [13.955392, 108.676758];

const ControllingGroup: React.FC = () => {
  const map = useMapEvents({
    overlayadd(e) {
      const bounds = new L.LatLngBounds(center, center);

      map.eachLayer(function (layer) {
        if (layer instanceof L.FeatureGroup) {
          bounds.extend(layer.getBounds());
        }
      });

      if (bounds.isValid()) {
        map.flyToBounds(bounds);
      }
    },
  });

  return null;
};

export default ControllingGroup;
