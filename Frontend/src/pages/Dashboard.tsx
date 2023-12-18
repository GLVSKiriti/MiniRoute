import "./Dashboard.scss";
import { useState } from "react";
import axios from "axios";
import { toast } from "react-toastify";
import MiniRoute from "../assets/MiniRoute.png";
import NavBar from "../components/NavBar";

function Dashboard() {
  const [miniRoute, setMiniRoute] = useState("");
  const [longurl, setLongUrl] = useState("");
  const [shortCode, setShortCode] = useState("");

  const shortenUrl = async () => {
    try {
      const token = sessionStorage.getItem("jwt");
      const res = await axios.post(
        `http://${location.hostname}:8080/url/shorten`,
        {
          longurl: longurl,
          shorturl: shortCode ? shortCode : null,
        },
        {
          headers: {
            Authorization: token,
          },
        }
      );

      setMiniRoute(`http://${location.hostname}:8080/url/redirect/` + res.data.shortUrl);
      toast.success("Successfully Shortened the URL", {
        position: "bottom-right",
      });
    } catch (error: any) {
      if (error.response) {
        toast.error(error.response.data, {
          position: "bottom-right",
        });
      }
    }
  };

  return (
    <>
      <div className="dashboardWrapper">
        <NavBar />
        <div className="dashboard">
          <div className="urlShortnerPad">
            <div className="inputTitle">
              🔗 Shorten a long URL<b>*</b>
            </div>
            <input
              type="text"
              placeholder="Enter long link here"
              value={longurl}
              onChange={(e) => setLongUrl(e.target.value)}
            />
            <div className="inputTitle">✨ Customize your link</div>
            <input
              type="text"
              placeholder="Enter custom short code"
              value={shortCode}
              onChange={(e) => setShortCode(e.target.value)}
            />
            <button disabled={!longurl} onClick={() => shortenUrl()}>
              Shorten URL
            </button>
            <div className="logo2">
              <img src={MiniRoute} alt="" />
            </div>
            <input
              placeholder="Your MiniRoute"
              className="miniroute"
              value={miniRoute}
              readOnly
            />
          </div>
        </div>
      </div>
    </>
  );
}

export default Dashboard;
