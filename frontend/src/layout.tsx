import { RouteSectionProps } from "@solidjs/router";
import { Suspense } from "solid-js";
import "~/index.css";

export default function Layout(props: RouteSectionProps) {
	return (
		<div>
			<a href="/">Index</a>
			<a href="/about">About nunu </a>
			<Suspense>{props.children}</Suspense>
		</div>
	)
}
