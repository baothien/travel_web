import React from "react";

import { Input, Space, Typography, AutoComplete } from "antd";
import Icon, {
  ClockCircleOutlined,
  EnvironmentOutlined,
  HomeOutlined,
  MenuOutlined,
  SearchOutlined,
} from "@ant-design/icons";
import { IconBaseProps } from "@ant-design/icons/lib/components/Icon";
import { Item } from "rc-menu";
import { useNavigate } from "react-router-dom";

const { Search } = Input;

const { Title, Text } = Typography;

const searchData = [
  {
    icon: <ClockCircleOutlined className="fs-5 text-secondary" />,
    title: "Tên địa điểm",
  },
  {
    icon: <EnvironmentOutlined className="fs-5 text-secondary" />,
    title: "Tên địa điểm",
  },
  {
    icon: <SearchOutlined className="fs-5 text-secondary" />,
    title: "Tên địa điểm",
  },
];

const renderItem = (title: string, icon: any) => ({
  value: title,
  label: (
    <Space style={{}} className="" size={10}>
      {icon}
      <Text className="text-secondary">{title}</Text>
    </Space>
  ),
});

const SearchBar: React.FC = () => {
  const navigate = useNavigate();

  const onSearch = (value: string) => console.log(value);

  const options = searchData.map((item) => renderItem(item.title, item.icon));

  return (
    <div style={{}} className="bg-white search-bar rounded">
      {/* <AutoComplete
        style={{ width: 350 }}
        popupClassName="certain-category-search-dropdown"
        dropdownMatchSelectWidth={350}
        options={options}
      >
        <Search
          style={{ width: 350 }}
          addonBefore={<MenuOutlined />}
          placeholder="Tìm kiếm"
          allowClear
          onSearch={onSearch}
        />
      </AutoComplete> */}

      <HomeOutlined
        style={{ fontSize: "2rem", padding: "0.5rem" }}
        onClick={() => navigate("/")}
      />
    </div>
  );
};

export default SearchBar;
