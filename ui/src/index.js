import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import Root from "./routes/root";
import ErrorPage from "./error-page";
import {
  List as ListListenerPage,
  New as NewListenerPage,
  Show as ShowListenerPage,
  Loader as ListenerLoader,
  Actions as ListenerActions
} from "./routes/listeners"

import {
  List as ListTargetGroupsPage,
  New as NewTargetGroupPage,
  Show as ShowTargetGroupPage,
  Loader as TargetGroupLoader,
  Actions as TargetGroupActions,
} from "./routes/target_groups"

import '@blueprintjs/core/lib/css/blueprint.css'

const router = createBrowserRouter([
  { path: "/", element: <Root />, errorElement: <ErrorPage />, children: [
      { path: "/listeners", children: [
          { path: "", element: <ListListenerPage />, loader: ListenerLoader.getListeners },
          { path: "/listeners/new", element: <NewListenerPage />, action: ListenerActions.createListener },
          { path: "/listeners/:listenerId", element: <ShowListenerPage />, loader: ListenerLoader.getListener },
        ]},
      { path: "/certificates", children: [
          { path: "" },
          { path: "/certificates/:certificateId" },
        ]},
      { path: "/target_groups", children: [
          { path: "", element: <ListTargetGroupsPage />, loader: TargetGroupLoader.getTargetGroups },
          { path: "/target_groups/new", element: <NewTargetGroupPage />, action: TargetGroupActions.createTargetGroup },
          { path: "/target_groups/:targetGroupId", element: <ShowTargetGroupPage />, loader: TargetGroupLoader.getTargetGroup, action: TargetGroupActions.addAttachment },
        ]}
    ]},
], {
  basename: "/ui/"
})

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
