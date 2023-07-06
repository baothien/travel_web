import { Button, Col, Form, Input, Row, Space, Typography } from "antd";
import React, { useState } from "react";
import profileApi from "../../apis/profileApi";
import { toast } from "react-toastify";

const { Text, Title } = Typography;

const ChangePassword: React.FC = () => {
  const [conditionPass, setConditionPass] = useState<any>({
    haveUppercase: false,
    haveNumber: false,
    haveFullLength: false,
  });

  const [passMatching, setPassMatching] = useState<boolean>(false);

  const [changePwForm] = Form.useForm<{ newPassword: string }>();
  const newPass = Form.useWatch("newPassword", changePwForm);

  const onFinish = (values: any) => {
    const params = {
      old_password: values.oldPassword,
      new_password: values.newPassword,
    };

    if (
      conditionPass.haveUppercase &&
      conditionPass.haveNumber &&
      conditionPass.haveFullLength
    ) {
      profileApi
        .changePassword(params)
        .then((res) => {
          toast.success(res.message);
        })
        .catch((err) => {
          toast.error(err.response.data.error_info.message[0]);
        });
    } else {
      toast.info("Chưa đủ điều kiện để đổi mật khẩu");
    }
  };

  function containsNumbers(str) {
    return /\d/.test(str);
  }

  function containsUppercase(str) {
    return /[A-Z]/.test(str);
  }

  const onChangeNewPass = (e) => {
    const value = e.target.value;

    //check length input
    if (value.length >= 8)
      setConditionPass((prev) => ({ ...prev, haveFullLength: true }));
    else setConditionPass((prev) => ({ ...prev, haveFullLength: false }));

    //check have number
    if (containsNumbers(value))
      setConditionPass((prev) => ({ ...prev, haveNumber: true }));
    else setConditionPass((prev) => ({ ...prev, haveNumber: false }));

    //check have uppercase
    if (containsUppercase(value))
      setConditionPass((prev) => ({ ...prev, haveUppercase: true }));
    else setConditionPass((prev) => ({ ...prev, haveUppercase: false }));
  };

  const onChangeConfPass = (e) => {
    if (e.target.value == newPass.toString()) {
      setPassMatching(true);
    } else {
      setPassMatching(false);
    }
  };

  return (
    <Row justify="space-between" align="middle">
      <Col span={12}>
        {" "}
        <Form
          name="changePwForm"
          form={changePwForm}
          className="change-pw-form"
          layout="vertical"
          onFinish={onFinish}
        >
          <Form.Item
            name="oldPassword"
            label="Mật khẩu cũ"
            rules={[{ required: true, message: "Không được để trống!" }]}
          >
            <Input.Password placeholder="Nhập mật khẩu cũ" />
          </Form.Item>

          <Form.Item
            name="newPassword"
            label="Mật khẩu mới"
            rules={[{ required: true, message: "Không được để trống!" }]}
          >
            <Input.Password
              style={{
                border:
                  conditionPass.haveFullLength &&
                  conditionPass.haveNumber &&
                  conditionPass.haveUppercase
                    ? "2px solid #12D47B"
                    : "",
              }}
              type="password"
              placeholder="Password"
              onChange={onChangeNewPass}
            />
          </Form.Item>

          <Form.Item
            name="confPassword"
            label="Xác nhận mật khẩu"
            dependencies={["newPassword"]}
            rules={[
              { required: true, message: "Không được để trống" },
              ({ getFieldValue }) => ({
                validator(rule, value) {
                  if (!value || getFieldValue("newPassword") === value) {
                    return Promise.resolve();
                  }
                  return Promise.reject("Mật khẩu chưa trùng nhau");
                },
              }),
            ]}
          >
            <Input.Password
              style={{
                border: passMatching ? "2px solid #12D47B" : "",
              }}
              type="password"
              placeholder="Password"
              onChange={onChangeConfPass}
            />
          </Form.Item>

          <Form.Item>
            <Button
              style={{ width: "100%", height: "2.2rem" }}
              className="login-form-button mt-1"
              type="primary"
              htmlType="submit"
            >
              Xác nhận
            </Button>
          </Form.Item>
        </Form>
      </Col>

      <Col span={9}>
        <Space direction="vertical">
          <Text style={{ fontWeight: "700", fontSize: "1.2rem" }}>
            Mật khẩu mới phải chứa
          </Text>
          <Text
            style={{
              color: conditionPass.haveUppercase ? "#12D47B" : "black",
            }}
            delete={conditionPass.haveUppercase}
            strong
          >
            Ít nhất 1 ký tự viết hoa (A-Z)
          </Text>
          <Text
            style={{
              color: conditionPass.haveNumber ? "#12D47B" : "black",
            }}
            delete={conditionPass.haveNumber}
            strong
          >
            Ít nhất 1 chữ số
          </Text>
          <Text
            style={{
              color: conditionPass.haveFullLength ? "#12D47B" : "black",
            }}
            delete={conditionPass.haveFullLength}
            strong
          >
            Ít nhất 8 ký tự
          </Text>
        </Space>
      </Col>
    </Row>
  );
};

export default ChangePassword;
