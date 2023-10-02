import "./Dashboard.scss";
import { useEffect } from "react";
import axios from "axios";
import { toast } from "react-toastify";

function Dashboard() {
  const fetchmyapi = async () => {
    try {
      const token = sessionStorage.getItem("jwt");
      const res = await axios.get("http://localhost:8080/url/shorten", {
        headers: {
          Authorization: token,
        },
      });
    } catch (error: any) {
      if (error.response) {
        toast.error(error.response.data.Message, {
          position: "bottom-right",
        });
      }
    }
  };

  useEffect(() => {
    fetchmyapi();
  }, []);

  return <div>This is dashboard</div>;
}

export default Dashboard;
