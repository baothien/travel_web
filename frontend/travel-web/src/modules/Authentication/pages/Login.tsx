import React from "react";
import { Link, useNavigate } from "react-router-dom";

import "./Authentication.scss";
import { Button, Col, Form, Input, Row, Typography } from "antd";
import {
  ArrowLeftOutlined,
  LockOutlined,
  MailOutlined,
} from "@ant-design/icons";

import logo from "../../../assets/logo/logo.png";
import mainLogo from "../../../assets/logo/mainLogo.jpg";
import { toast } from "react-toastify";
import { authActions } from "../authSlice";
import authApi from "../../../apis/authApi";
import { LoginParams } from "../../../models/common";
import { useAppDispatch } from "../../../hooks";

const { Title, Text } = Typography;

const Login: React.FC = () => {
  const dispatch = useAppDispatch();

  const navigate = useNavigate();

  const onFinish = async (value: any) => {
    const params: LoginParams = {
      user_name: value.email,
      password: value.password,
    };

    await authApi
      .postLogin(params)
      .then((res) => {
        const user = res.data;

        dispatch(authActions.loginSuccess(user));

        toast.success(`Xin chào ${user.user_name || ""}`);
        navigate(-1);
      })
      .catch((err: any) => {
        toast.error("Đăng nhập thất bại");
        console.log("error login", err);
      });
  };

  return (
    <div className="text-white primary-font">
      <div className="bg-image"></div>

      <Button
        className="bg-btn d-flex align-items-center justify-content-center"
        type="text"
        icon={<ArrowLeftOutlined className="text-dark" />}
        onClick={() => navigate(-1)}
      >
        <Text className="text-dark fs-6">Quay lại</Text>
      </Button>

      <div className="bg-text text-white rounded">
        <Row justify="center">
          <Text
            style={{ fontSize: "1.1rem" }}
            className="text-white-50 fw-light mb-4"
          >
            Đăng nhập để sử dụng dịch vụ
          </Text>
        </Row>

        <Row className="" justify="center">
          <Form className="login-form" onFinish={onFinish}>
            <Form.Item className="" name="email">
              <Input
                style={{
                  backgroundColor: "rgba(0,0,0, 0.4)",
                  color: "#fff",
                  padding: "0.5rem",
                  width: "13rem",
                }}
                className="border-0 rounded m-0"
                placeholder="Nhập email"
                prefix={<MailOutlined />}
                defaultValue=""
                allowClear
              />
            </Form.Item>

            <Form.Item className="" name="password">
              <Input.Password
                style={{
                  backgroundColor: "rgba(0,0,0, 0.4)",
                  color: "#fff",
                  padding: "0.5rem",
                  width: "13rem",
                }}
                className="border-0 rounded"
                placeholder="Nhập mật khẩu"
                prefix={<LockOutlined />}
              />
            </Form.Item>

            <Form.Item>
              <Button
                style={{
                  backgroundColor: "#FD7E14",
                  width: "13rem",
                  height: "2.2rem",
                }}
                className="border-0 text-white rounded"
                htmlType="submit"
              >
                Đăng nhập
              </Button>
            </Form.Item>
          </Form>
        </Row>

        <Row justify="center">
          <Button
            style={{ fontSize: "0.8rem" }}
            className="text-white"
            type="text"
          >
            Quên mật khẩu?
          </Button>
        </Row>

        <Row justify="center">
          <Text style={{ fontSize: "0.8rem" }} className="text-white">
            Chưa có tài khoản?
            <Link to={"/register"}>
              <span style={{ color: "#FD7E14" }}> Đăng ký</span>
            </Link>
          </Text>
        </Row>
      </div>
    </div>
  );
};

export default Login;
