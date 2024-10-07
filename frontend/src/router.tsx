import { Route, Router as SolidRouter } from "@solidjs/router";
import Layout from "~/layout";
import NotFound from "~/routes/404";
import Home from "~/routes/home";

export default function Router() {
	return (
		<SolidRouter root={Layout}>
			<Route path="/" component={Home} />
			<Route path="*404" component={NotFound} />
		</SolidRouter>
	)
}
