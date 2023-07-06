import React, { useEffect, useRef, useState } from "react";
import { PlayCircleOutlined } from "@ant-design/icons";
import { Modal } from "antd";

const Video: React.FC = () => {
  const [state, setState] = useState(false);

  const showModal = () => {
    setState(true);
  };

  const hideModal = () => {
    setState(false);
  };
  return (
    <div>
      <img
        width={480}
        className="video-detail"
        height={500}
        alt=""
        src="https://backpacktraveler.qodeinteractive.com/wp-content/uploads/2018/08/brazil-single-2-2.jpg"
      />
      <PlayCircleOutlined
        style={{ fontSize: "4rem", top: "43%", right: "43%" }}
        className="position-absolute text-white"
        onClick={showModal}
      />
      <Modal
        title="Video"
        open={state}
        footer={null}
        onCancel={hideModal}
        // afterClose={this.pause}
        bodyStyle={{ padding: 0 }}
      >
        add video vào nè
      </Modal>
    </div>
  );
};

export default Video;
