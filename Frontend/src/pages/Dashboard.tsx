import "./Dashboard.scss";
import { useState } from "react";
import MiniRoute from "../assets/MiniRoute.png";

function Dashboard() {
  const [miniRoute, setMiniRoute] = useState("");

  return (
    <>
      <div className="dashboardWrapper">
        <div className="navBar">
          <div className="logo">
            <img src={MiniRoute} alt="" />
          </div>
          <div className="menu">
            <div className="myurls">My Urls</div>
            <div className="logout">LogOut</div>
          </div>
        </div>
        <div className="dashboard">
          <div className="urlShortnerPad">
            <div className="inputTitle">
              🔗 Shorten a long URL<b>*</b>
            </div>
            <input type="text" placeholder="Enter long link here" />
            <div className="inputTitle">✨ Customize your link</div>
            <input type="text" placeholder="Enter custom short code" />
            <button>Shorten URL</button>
            <div className="logo2">
              <img src={MiniRoute} alt="" />
            </div>
            <input
              type="text"
              value={miniRoute}
              contentEditable={false}
              placeholder="Your MiniRoute"
              className="miniroute"
            />
          </div>
        </div>
      </div>
    </>
  );
}

export default Dashboard;
