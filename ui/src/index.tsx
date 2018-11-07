import * as React from "react"; 
import * as ReactDOM from "react-dom";
import {HashRouter as Router, Route} from "react-router-dom";


import {EventListContainer} from "./components/event_list_container";
import {Navigation} from "./components/navigation";


class App extends React.Component<{}, {}>{ 
	render() {
		const eventList = () =><EventListContainer eventServiceURL="http://localhost:8181"/>

        return <Router>
            <div>
                <Navigation brandName="MyEvents"/>
                <div className="container">
                    <h1>My Events</h1>
                    <Route exact path="/" component={eventList}/>
                </div>
            </div>
        </Router>
	} 
} 

ReactDOM.render( 
	<App/>,
	document.getElementById("myevents-app")
);