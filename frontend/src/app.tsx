import { MetaProvider } from "@solidjs/meta";
import { Route, Router } from "@solidjs/router";
import { FileRoutes } from "@solidjs/start/router";
import { Suspense } from "solid-js";

import "./index.css";
import Layout from "./layout";
import Home from "./routes/home";
import NotFound from "./routes/404";

export default function App() {

  return (
    <MetaProvider>
      <Router root={Layout}>
        <Route path="/" component={Home} />
        <Route path="*404" component={NotFound} />
      </Router>
    </MetaProvider>
  );
}
