import { styled } from "solid-styled-components";

export const Button = styled("button")`
	font-family: inherit;
	font-size: inherit;
	padding: 1em 2em;
	color: #335d92;
	background-color: rgba(30, 107, 158, 0.1);
	border-radius: 2em;
	border: 2px solid rgba(68, 107, 158, 0);
	outline: none;
	width: 200px;
	font-variant-numeric: tabular-nums;
	&:focus {
		border: 2px solid #335d92;
	}
`
