/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, { enumerable: true, get: getter });
/******/ 		}
/******/ 	};
/******/
/******/ 	// define __esModule on exports
/******/ 	__webpack_require__.r = function(exports) {
/******/ 		if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 			Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 		}
/******/ 		Object.defineProperty(exports, '__esModule', { value: true });
/******/ 	};
/******/
/******/ 	// create a fake namespace object
/******/ 	// mode & 1: value is a module id, require it
/******/ 	// mode & 2: merge all properties of value into the ns
/******/ 	// mode & 4: return value when already ns object
/******/ 	// mode & 8|1: behave like require
/******/ 	__webpack_require__.t = function(value, mode) {
/******/ 		if(mode & 1) value = __webpack_require__(value);
/******/ 		if(mode & 8) return value;
/******/ 		if((mode & 4) && typeof value === 'object' && value && value.__esModule) return value;
/******/ 		var ns = Object.create(null);
/******/ 		__webpack_require__.r(ns);
/******/ 		Object.defineProperty(ns, 'default', { enumerable: true, value: value });
/******/ 		if(mode & 2 && typeof value != 'string') for(var key in value) __webpack_require__.d(ns, key, function(key) { return value[key]; }.bind(null, key));
/******/ 		return ns;
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = "./src/index.tsx");
/******/ })
/************************************************************************/
/******/ ({

/***/ "./src/components/event_list.tsx":
/*!***************************************!*\
  !*** ./src/components/event_list.tsx ***!
  \***************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
eval("\nvar __extends = (this && this.__extends) || (function () {\n    var extendStatics = function (d, b) {\n        extendStatics = Object.setPrototypeOf ||\n            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||\n            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };\n        return extendStatics(d, b);\n    }\n    return function (d, b) {\n        extendStatics(d, b);\n        function __() { this.constructor = d; }\n        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());\n    };\n})();\nObject.defineProperty(exports, \"__esModule\", { value: true });\nvar event_list_item_1 = __webpack_require__(/*! ./event_list_item */ \"./src/components/event_list_item.tsx\");\nvar React = __webpack_require__(/*! react */ \"react\");\nvar EventList = /** @class */ (function (_super) {\n    __extends(EventList, _super);\n    function EventList() {\n        return _super !== null && _super.apply(this, arguments) || this;\n    }\n    EventList.prototype.render = function () {\n        var items = this.props.events.map(function (e) {\n            return React.createElement(event_list_item_1.EventListItem, { event: e });\n        });\n        return React.createElement(\"table\", { className: \"table\" },\n            React.createElement(\"thead\", null,\n                React.createElement(\"tr\", null,\n                    React.createElement(\"th\", null, \"Event\"),\n                    React.createElement(\"th\", null, \"Where\"),\n                    React.createElement(\"th\", { colSpan: 2 }, \"When (start/end)\"),\n                    React.createElement(\"th\", null, \"Actions\"))),\n            React.createElement(\"tbody\", null, items));\n    };\n    return EventList;\n}(React.Component));\nexports.EventList = EventList;\n\n\n//# sourceURL=webpack:///./src/components/event_list.tsx?");

/***/ }),

/***/ "./src/components/event_list_container.tsx":
/*!*************************************************!*\
  !*** ./src/components/event_list_container.tsx ***!
  \*************************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
eval("\nvar __extends = (this && this.__extends) || (function () {\n    var extendStatics = function (d, b) {\n        extendStatics = Object.setPrototypeOf ||\n            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||\n            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };\n        return extendStatics(d, b);\n    }\n    return function (d, b) {\n        extendStatics(d, b);\n        function __() { this.constructor = d; }\n        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());\n    };\n})();\nObject.defineProperty(exports, \"__esModule\", { value: true });\nvar React = __webpack_require__(/*! react */ \"react\");\nvar event_list_1 = __webpack_require__(/*! ./event_list */ \"./src/components/event_list.tsx\");\nvar EventListContainer = /** @class */ (function (_super) {\n    __extends(EventListContainer, _super);\n    function EventListContainer(p) {\n        var _this = _super.call(this, p) || this;\n        _this.state = {\n            loading: true,\n            events: []\n        };\n        console.log(\"getting events from: [\" + p.eventListURL + \"/events\" + \"]\");\n        fetch(p.eventListURL + \"/events\", { method: \"GET\" })\n            .then(function (response) { return response.json(); })\n            .then(function (events) {\n            _this.setState({\n                loading: false,\n                events: events\n            });\n        });\n        return _this;\n    }\n    EventListContainer.prototype.render = function () {\n        if (this.state.loading) {\n            return React.createElement(\"div\", null, \"Loading...\");\n        }\n        return React.createElement(event_list_1.EventList, { events: this.state.events });\n    };\n    return EventListContainer;\n}(React.Component));\nexports.EventListContainer = EventListContainer;\n\n\n//# sourceURL=webpack:///./src/components/event_list_container.tsx?");

/***/ }),

/***/ "./src/components/event_list_item.tsx":
/*!********************************************!*\
  !*** ./src/components/event_list_item.tsx ***!
  \********************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
eval("\nvar __extends = (this && this.__extends) || (function () {\n    var extendStatics = function (d, b) {\n        extendStatics = Object.setPrototypeOf ||\n            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||\n            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };\n        return extendStatics(d, b);\n    }\n    return function (d, b) {\n        extendStatics(d, b);\n        function __() { this.constructor = d; }\n        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());\n    };\n})();\nObject.defineProperty(exports, \"__esModule\", { value: true });\nvar React = __webpack_require__(/*! react */ \"react\");\nvar EventListItem = /** @class */ (function (_super) {\n    __extends(EventListItem, _super);\n    function EventListItem() {\n        return _super !== null && _super.apply(this, arguments) || this;\n    }\n    EventListItem.prototype.render = function () {\n        var start = new Date(this.props.event.StartDate * 1000);\n        var end = new Date(this.props.event.EndDate * 1000);\n        return React.createElement(\"tr\", null,\n            React.createElement(\"td\", null, this.props.event.Name),\n            React.createElement(\"td\", null, this.props.event.Location.Name),\n            React.createElement(\"td\", null, start.toLocaleDateString()),\n            React.createElement(\"td\", null, end.toLocaleDateString()),\n            React.createElement(\"td\", null));\n    };\n    return EventListItem;\n}(React.Component));\nexports.EventListItem = EventListItem;\n\n\n//# sourceURL=webpack:///./src/components/event_list_item.tsx?");

/***/ }),

/***/ "./src/index.tsx":
/*!***********************!*\
  !*** ./src/index.tsx ***!
  \***********************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
eval("\nvar __extends = (this && this.__extends) || (function () {\n    var extendStatics = function (d, b) {\n        extendStatics = Object.setPrototypeOf ||\n            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||\n            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };\n        return extendStatics(d, b);\n    }\n    return function (d, b) {\n        extendStatics(d, b);\n        function __() { this.constructor = d; }\n        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());\n    };\n})();\nObject.defineProperty(exports, \"__esModule\", { value: true });\nvar React = __webpack_require__(/*! react */ \"react\");\nvar ReactDOM = __webpack_require__(/*! react-dom */ \"react-dom\");\nvar event_list_container_1 = __webpack_require__(/*! ./components/event_list_container */ \"./src/components/event_list_container.tsx\");\nvar App = /** @class */ (function (_super) {\n    __extends(App, _super);\n    function App() {\n        return _super !== null && _super.apply(this, arguments) || this;\n    }\n    App.prototype.render = function () {\n        return React.createElement(\"div\", { className: \"container\" },\n            React.createElement(\"h1\", null, \"MyEvents\"),\n            React.createElement(event_list_container_1.EventListContainer, { eventListURL: \"http://localhost:8181\" }));\n    };\n    return App;\n}(React.Component));\nReactDOM.render(React.createElement(App, null), document.getElementById(\"myevents-app\"));\n\n\n//# sourceURL=webpack:///./src/index.tsx?");

/***/ }),

/***/ "react":
/*!************************!*\
  !*** external "React" ***!
  \************************/
/*! no static exports found */
/***/ (function(module, exports) {

eval("module.exports = React;\n\n//# sourceURL=webpack:///external_%22React%22?");

/***/ }),

/***/ "react-dom":
/*!***************************!*\
  !*** external "ReactDOM" ***!
  \***************************/
/*! no static exports found */
/***/ (function(module, exports) {

eval("module.exports = ReactDOM;\n\n//# sourceURL=webpack:///external_%22ReactDOM%22?");

/***/ })

/******/ });