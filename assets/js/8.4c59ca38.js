(window.webpackJsonp=window.webpackJsonp||[]).push([[8],{328:function(t,a,s){"use strict";s.r(a);var e=s(33),n=Object(e.a)({},(function(){var t=this,a=t.$createElement,s=t._self._c||a;return s("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[s("h1",{attrs:{id:"getting-started"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#getting-started"}},[t._v("#")]),t._v(" Getting Started")]),t._v(" "),s("p",[t._v("The quick start creates a Hello World service using "),s("a",{attrs:{href:"https://github.com/bongnv/gokit",target:"_blank",rel:"noopener noreferrer"}},[s("code",[t._v("gokitgen")]),s("OutboundLink")],1),t._v(".")]),t._v(" "),s("h2",{attrs:{id:"install-gokitgen"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#install-gokitgen"}},[t._v("#")]),t._v(" Install "),s("code",[t._v("gokitgen")])]),t._v(" "),s("p",[t._v("Make sure "),s("a",{attrs:{href:"https://golang.org/doc/install",target:"_blank",rel:"noopener noreferrer"}},[t._v("Golang"),s("OutboundLink")],1),t._v(" is installed. Use the following command to install "),s("code",[t._v("gokitgen")]),t._v(" from source:")]),t._v(" "),s("div",{staticClass:"language-bash extra-class"},[s("pre",{pre:!0,attrs:{class:"language-bash"}},[s("code",[t._v("go get -u github.com/bongnv/gokit/cmd/gokitgen\n")])])]),s("h2",{attrs:{id:"scaffold-a-project"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#scaffold-a-project"}},[t._v("#")]),t._v(" Scaffold a project")]),t._v(" "),s("div",{staticClass:"language-bash extra-class"},[s("pre",{pre:!0,attrs:{class:"language-bash"}},[s("code",[s("span",{pre:!0,attrs:{class:"token function"}},[t._v("mkdir")]),t._v(" hello\n"),s("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("cd")]),t._v(" hello\n")])])]),s("p",[t._v("Use "),s("code",[t._v("gokitgen")]),t._v(" to scaffold an empty project:")]),t._v(" "),s("div",{staticClass:"language-bash extra-class"},[s("pre",{pre:!0,attrs:{class:"language-bash"}},[s("code",[t._v("gokitgen scaffold -package github.com/hello\n")])])]),s("p",[t._v("By default, "),s("code",[t._v("service.go")]),t._v(" will be generated with two example endpoints.")]),t._v(" "),s("div",{staticClass:"language-go extra-class"},[s("pre",{pre:!0,attrs:{class:"language-go"}},[s("code",[s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("//go:generate gokitgen service -interface Service")]),t._v("\n"),s("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("type")]),t._v(" Service "),s("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("interface")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\t"),s("span",{pre:!0,attrs:{class:"token function"}},[t._v("Hello")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("ctx context"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Context"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" p "),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("Request"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("Response"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("error")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\t"),s("span",{pre:!0,attrs:{class:"token function"}},[t._v("Bye")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("ctx context"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Context"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" req "),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("ByeRequest"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("ByeResponse"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("error")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])])])}),[],!1,null,null,null);a.default=n.exports}}]);