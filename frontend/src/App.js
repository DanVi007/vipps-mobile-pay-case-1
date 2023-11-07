import { useState } from "react";
import axios from "axios";
import './App.css'

function App() {
    const [topicCount, setTopicCount] = useState(0);
    const [topic, setTopic] = useState("");

    const backendApi = axios.create({
        baseURL: "http://localhost:8080"
    });

    const fetchTopicCountFromApi = async () => {
        try {
            const response = await backendApi.get("/topic_counter?topic=" + topic);
            if (response.data !== null) {
                setTopicCount(response.data)
            } else {
                console.error("Failed to fetch data from the API");
            }
        } catch (error) {
            console.error("An error occurred while fetching data: ", error);
        }

    }

    const handleSubmit = (event) => {
        event.preventDefault();
        if (topic !== "" && topic.length !== 0 && /\S/.test(topic)) {
            fetchTopicCountFromApi();
        }
    }

    return (
        <div className="mainPage" >
            <h1>Vipps MobilePay</h1>
            <div className="container">
                <div className="content" >
                    <form id="topicForm" onSubmit={handleSubmit}>
                        <label for="topic"><b>Topic: </b></label>
                        <input type="text" id="topic" name="topic" value={topic} onChange={(e) => setTopic(e.target.value)} />
                    </form>
                    <p><b>Amount of times topic is mentioned in article: </b>{topicCount}</p>
                </div>
            </div>
        </div>
    )
}

export default App;