import "./MyUrls.scss";
import NavBar from "../components/NavBar";
import { toast } from "react-toastify";
import axios from "axios";
import { useState, useEffect } from "react";
import { Table } from "antd";
import CopyIcon from "../components/CopyIcon";

type tableRow = {
  miniroute: string;
  longurl: string;
};

function MyUrls() {
  const [dataSource, setDataSource] = useState<tableRow[]>([]);
  const [hoveredRowIndex, setHoveredRowIndex] = useState(-1);
  const [hoveredRowIndex2, setHoveredRowIndex2] = useState(-1);

  const columns: any = [
    {
      title: "MiniRoute",
      dataIndex: "shortUrl",
      onCell: (_: any, rowIndex: number) => ({
        onMouseEnter: () => setHoveredRowIndex(rowIndex),
        onMouseLeave: () => setHoveredRowIndex(-1),
      }),
      render: (value: string, _: any, rowIndex: number) => {
        return (
          <span style={{ position: "relative" }}>
            http://localhost:8080/url/redirect/{value}
            <CopyIcon
              visibility={hoveredRowIndex === rowIndex ? "visible" : "hidden"}
              copyText={`http://localhost:8080/url/redirect/${value}`}
            />
          </span>
        );
      },
    },
    {
      title: "Original URL",
      dataIndex: "longUrl",
      onCell: (_: any, rowIndex: number) => ({
        onMouseEnter: () => setHoveredRowIndex2(rowIndex),
        onMouseLeave: () => setHoveredRowIndex2(-1),
      }),
      render: (value: string, _: string, rowIndex: number) => {
        return (
          <span style={{ position: "relative" }}>
            {value}
            <CopyIcon
              visibility={hoveredRowIndex2 === rowIndex ? "visible" : "hidden"}
              copyText={value}
            />
          </span>
        );
      },
    },
  ];
  const fetchMyUrls = async () => {
    try {
      const token = sessionStorage.getItem("jwt");
      const res = await axios.get("http://localhost:8080/url/myurls", {
        headers: {
          Authorization: token,
        },
      });
      setDataSource(res.data);
    } catch (error: any) {
      if (error.response) {
        toast.error(error.response.data, {
          position: "bottom-right",
        });
      }
    }
  };

  useEffect(() => {
    fetchMyUrls();
  }, []);

  return (
    <>
      <div className="MyUrlsWrapper">
        <NavBar />
        <div className="myurlslist">
          <Table
            className="urltable"
            dataSource={dataSource}
            columns={columns}
            pagination={{ pageSize: 8, position: ["bottomCenter"] }}
          />
        </div>
      </div>
    </>
  );
}

export default MyUrls;
