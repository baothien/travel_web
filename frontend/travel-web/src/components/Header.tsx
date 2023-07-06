import React from "react";

import { Link } from "react-router-dom";

import {
  Typography,
  Row,
  Col,
  Space,
  Button,
  MenuProps,
  Dropdown,
  message,
} from "antd";
import { useNavigate } from "react-router-dom";

import locationIcon from "../assets/logo/locationicon.jpg";

import logo1 from "../assets/logo/1.png";
import logo2 from "../assets/logo/2.png";
import logo3 from "../assets/logo/3.png";
import mainLogo from "../assets/logo/mainLogo.jpg";

import GlobalLogo from "../assets/global-logo.png";

import { UserOutlined } from "@ant-design/icons";

import { selectUser } from "../modules/Authentication/authSlice";
import { useAppDispatch, useAppSelector } from "../hooks";
import { NotifyHeader } from "../modules/Notification/NotifyHeader";
import { notifyActions } from "../modules/Notification/notifySlice";

const { Text } = Typography;

const items: MenuProps["items"] = [
  {
    label: <a href="https://www.antgroup.com">Tiếng Việt</a>,
    key: "0",
  },
  {
    label: <a href="https://www.aliyun.com">Tiếng Anh</a>,
    key: "1",
  },
];

const Header: React.FC = () => {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();

  const user = useAppSelector(selectUser);

  if (user) {
    const ws = new WebSocket(
      `wss://travel-api.huytx.com/stag/notify-service/ws/${user?.id}`
    );

    React.useEffect(() => {
      ws.onmessage = function (event: any) {
        message.info(JSON.parse(event.data).title);
        dispatch(notifyActions.increaseCount());
      };
    });
  }

  return (
    <Row className="header-homepage" justify="space-between" align="middle">
      {/* <Col>
        <Dropdown menu={{ items }} trigger={["click"]}>
          <a
            style={{ textDecoration: "none" }}
            onClick={(e) => e.preventDefault()}
          >
            <Space>
              <img width={25} src={GlobalLogo} alt="logo language" />
              <Text>Ngôn ngữ</Text>
            </Space>
          </a>
        </Dropdown>
      </Col> */}

      <Col>
        <Row justify="center" align="middle">
          <Col>
            <Link to={"/"}>
              {/* <img height={80} width={110} src={locationIcon} alt="" /> */}
              {/* <img height={80} width={110} src={logo3} alt="" /> */}
              <img height={100} src={mainLogo} />
            </Link>
          </Col>

          <Button size="small" type="text" href="/">
            TRANG CHỦ
          </Button>
          <Button
            size="small"
            type="text"
            href="/introduce"
            hrefLang="/introduce"
          >
            GIỚI THIỆU
          </Button>

          {/* <Button size="small" type="text">
            TIN TỨC
          </Button>
          <Button size="small" type="text">
            LIÊN HỆ
          </Button> */}
        </Row>
      </Col>

      <Col>
        <Row justify="center" align="middle">
          {user ? (
            <>
              <NotifyHeader />

              <Button
                className="d-flex align-items-center"
                type="text"
                icon={<UserOutlined style={{ fontSize: "1.2rem" }} />}
                onClick={() => navigate("/profile/user-information")}
              >
                {/* {user.user_name} */}
                Tài khoản
              </Button>
            </>
          ) : (
            <>
              <Link to={"/login"}>
                <Button size="small" type="text">
                  ĐĂNG NHẬP
                </Button>
              </Link>

              <Link to={"/register"}>
                <Button size="small" type="text">
                  ĐĂNG KÝ
                </Button>
              </Link>
            </>
          )}
        </Row>
      </Col>
    </Row>
  );
};

export default Header;
