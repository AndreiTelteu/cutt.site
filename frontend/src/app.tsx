import { MetaProvider } from "@solidjs/meta";

import Router from "~/router";

export default function App() {
	return (
		<MetaProvider>
			<Router />
		</MetaProvider>
	);
}
