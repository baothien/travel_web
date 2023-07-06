import React from "react";
import { LayersControl, FeatureGroup, Polygon } from "react-leaflet";
import { latLng, LatLng, LocationEvent } from "leaflet";

import { tileLayer } from "../../../utils";

import L from "leaflet";

const River: [number, number][] = [
  [52.261433597272294, 21.01272583007813],
  [52.259857564663236, 21.008262634277347],
  [52.2485084763669, 21.019763946533207],
  [52.24377883252849, 21.026287078857425],
  [52.240835688576325, 21.032123565673828],
  [52.235474460511696, 21.037960052490238],
  [52.22800971166361, 21.04465484619141],
  [52.2249603686772, 21.047916412353516],
  [52.23600009974023, 21.042766571044925],
  [52.24966453484508, 21.023368835449222],
  [52.26132853017426, 21.01272583007813],
];

const Place: [number, number][] = [
  [52.23295130556395, 20.998821258544925],
  [52.230007443518716, 20.98474502563477],
  [52.224750061441355, 20.98920822143555],
  [52.22769427210073, 21.002597808837894],
];

const riverColor = { color: "red" };
const placeColor = { color: "blue" };

const ControlPolygon: React.FC = () => {
  return (
    <LayersControl position="topright" collapsed={false}>
      <LayersControl.Overlay name="River">
        <FeatureGroup>
          <Polygon pathOptions={riverColor} positions={River} />
        </FeatureGroup>
      </LayersControl.Overlay>

      <LayersControl.Overlay name="Place">
        <FeatureGroup>
          <Polygon pathOptions={placeColor} positions={Place} />
        </FeatureGroup>
      </LayersControl.Overlay>
    </LayersControl>
  );
};

export default ControlPolygon;
