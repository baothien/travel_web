import React, { useRef, useState } from "react";

import {
  Avatar,
  Button,
  Form,
  Input,
  Row,
  Space,
  Typography,
  Upload,
  UploadFile,
} from "antd";

import { useAppSelector } from "../../../../hooks";
import { selectUser } from "../../../Authentication/authSlice";
import { ReplyReviewParams } from "../../../../models/review";
import reviewApi from "../../../../apis/reviewApi";
import { toast } from "react-toastify";
import { RcFile, UploadProps } from "antd/es/upload";
import uploadApi from "../../../../apis/uploadApi";
import { PlusOutlined } from "@ant-design/icons";
import dayjs from "dayjs";
import { useForm } from "antd/es/form/Form";

export interface IReplyActionProps {
  cmtId: string;
  reRender: React.Dispatch<React.SetStateAction<boolean>>;
}

const { Text } = Typography;

const { TextArea } = Input;

export function ChildReviewAction({ cmtId, reRender }: IReplyActionProps) {
  const [form] = useForm();

  const user = useAppSelector(selectUser);

  const today = dayjs().format("YYYY-MM-DD");

  const [previewOpen, setPreviewOpen] = useState(false);
  const [previewImage, setPreviewImage] = useState("");
  const [previewTitle, setPreviewTitle] = useState("");

  const [fileList, setFileList] = useState<UploadFile[]>([]);

  const handlePreview = async (file: UploadFile) => {
    if (!file.url && !file.preview) {
      file.preview = await getBase64(file.originFileObj as RcFile);
    }

    setPreviewImage(file.url || (file.preview as string));
    setPreviewOpen(true);
    setPreviewTitle(
      file.name || file.url!.substring(file.url!.lastIndexOf("/") + 1)
    );
  };

  const handleSubmitReply = (values: any) => {
    const params: ReplyReviewParams = {
      description: values.answer,
      parent_id: cmtId,
      review_img: fileList.map((item) => ({ name: item.name, url: item.url })),
    };

    reviewApi
      .postReplyReview(params)
      .then((res) => {
        toast.success("Phản hổi bình luận thành công");
        form.resetFields();

        reRender((prev) => !prev);
      })
      .catch((err) => {
        toast.error("Phản hồi thất bại");
        console.log(err);
      });
  };

  const handleChangeImage: UploadProps["onChange"] = (
    //   {
    //   fileList: newFileList,
    // }
    value: any
  ) => {
    const fileObj = value.file.originFileObj;
    uploadApi
      .upload({ file: fileObj, type: "review" })
      .then((res) => {
        const resImage = res.data;
        const file: UploadFile = {
          uid: resImage.id,
          name: resImage.name,
          status: "done",
          url: resImage.full_path,
        };
        setFileList((prev) => [...prev, file]);
      })
      .catch((err) => console.log(err));
  };

  return (
    <div
      style={{ padding: "0.8rem", border: "1px solid #cfcfcf" }}
      className="mt-2 rounded reply-action"
    >
      <Row className="w-100 mb-1" justify="space-between" align="middle">
        <Space>
          <Avatar size={40} src={user?.avatar} />

          <Text
            style={{ fontSize: "0.8rem" }}
            className="text-align-center"
            strong
          >
            {user?.full_name}
          </Text>
        </Space>

        <Text className="text-secondary">{today}</Text>
      </Row>

      <Form form={form} onFinish={handleSubmitReply}>
        <Form.Item
          style={{ marginBottom: "0.8rem" }}
          name="answer"
          rules={[{ required: true, message: "Hãy nhập nội dung phản hồi" }]}
        >
          <TextArea
            placeholder="Nhập bình luận"
            autoSize={{ minRows: 1, maxRows: 6 }}
          />
        </Form.Item>

        <Form.Item style={{ margin: "0" }}>
          <Upload
            action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
            listType="picture-card"
            fileList={fileList}
            onPreview={handlePreview}
            onChange={handleChangeImage}
          >
            {fileList.length >= 8 ? null : uploadButton}
          </Upload>
        </Form.Item>

        <Row className="mt-1" justify="end">
          <Form.Item style={{ margin: "0" }}>
            <Button
              style={{
                backgroundColor: "#69B9C7",
                width: "100%",
              }}
              htmlType="submit"
            >
              <Text className="text-white">Trả lời</Text>
            </Button>
          </Form.Item>
        </Row>
      </Form>
    </div>
  );
}

const getBase64 = (file: RcFile): Promise<string> =>
  new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result as string);
    reader.onerror = (error) => reject(error);
  });

const uploadButton = (
  <div>
    <PlusOutlined />
    <div style={{ marginTop: 8 }}>Upload</div>
  </div>
);
