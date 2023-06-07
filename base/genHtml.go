package base

func GenStartHtml() (result string) {
	result = `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<style>
		@font-face { font-family: "FZFangSong-Z02"; src:local("FZFangSong-Z02"), url("https://imgcdn.umiwi.com/ttf/fangzhengfangsong_gbk.ttf"); }
		@font-face { font-family: "FZKai-Z03"; src:local("FZKai-Z03"), url("https://imgcdn.umiwi.com/ttf/fangzhengkaiti_gbk.ttf"); }
		@font-face { font-family: "PingFang SC"; src:local("PingFang SC"); }
		@font-face { font-family: "DeDaoJinKai"; src:local("DeDaoJinKai"), url("https://imgcdn.umiwi.com/ttf/dedaojinkaiw03.ttf");}
		@font-face { font-family: "Source Code Pro"; src:local("Source Code Pro"), url("https://imgcdn.umiwi.com/ttf/0315911806889993935644188722660020367983.ttf"); }
		table, tr, td, th, tbody, thead, tfoot {page-break-inside: avoid !important;}
pre code.hljs {
    display: block;
    overflow-x: auto;
    padding: 1em
}

code.hljs {
    padding: 3px 5px
}

/*!
  Theme: GitHub
  Description: Light theme as seen on github.com
  Author: github.com
  Maintainer: @Hirse
  Updated: 2021-05-15

  Outdated base version: https://github.com/primer/github-syntax-light
  Current colors taken from GitHub's CSS
*/
.hljs {
    color: #24292e;
    background: #fff
}

.hljs-doctag,.hljs-keyword,.hljs-meta .hljs-keyword,.hljs-template-tag,.hljs-template-variable,.hljs-type,.hljs-variable.language_ {
    color: #d73a49
}

.hljs-title,.hljs-title.class_,.hljs-title.class_.inherited__,.hljs-title.function_ {
    color: #6f42c1
}

.hljs-attr,.hljs-attribute,.hljs-literal,.hljs-meta,.hljs-number,.hljs-operator,.hljs-selector-attr,.hljs-selector-class,.hljs-selector-id,.hljs-variable {
    color: #005cc5
}

.hljs-meta .hljs-string,.hljs-regexp,.hljs-string {
    color: #032f62
}

.hljs-built_in,.hljs-symbol {
    color: #e36209
}

.hljs-code,.hljs-comment,.hljs-formula {
    color: #6a737d
}

.hljs-name,.hljs-quote,.hljs-selector-pseudo,.hljs-selector-tag {
    color: #22863a
}

.hljs-subst {
    color: #24292e
}

.hljs-section {
    color: #005cc5;
    font-weight: 700
}

.hljs-bullet {
    color: #735c0f
}

.hljs-emphasis {
    color: #24292e;
    font-style: italic
}

.hljs-strong {
    color: #24292e;
    font-weight: 700
}

.hljs-addition {
    color: #22863a;
    background-color: #f0fff4
}

.hljs-deletion {
    color: #b31d28;
    background-color: #ffeef0
}

.flatpickr-calendar {
    background: transparent;
    opacity: 0;
    display: none;
    text-align: center;
    visibility: hidden;
    padding: 0;
    animation: none;
    direction: ltr;
    border: 0;
    font-size: 14px;
    line-height: 24px;
    border-radius: 5px;
    position: absolute;
    width: 307.875px;
    box-sizing: border-box;
    touch-action: manipulation;
    background: #fff;
    box-shadow: 1px 0 #e6e6e6,-1px 0 #e6e6e6,0 1px #e6e6e6,0 -1px #e6e6e6,0 3px 13px #00000014
}

.flatpickr-calendar.open,.flatpickr-calendar.inline {
    opacity: 1;
    max-height: 640px;
    visibility: visible
}

.flatpickr-calendar.open {
    display: inline-block;
    z-index: 99999
}

.flatpickr-calendar.animate.open {
    animation: fpFadeInDown .3s cubic-bezier(.23,1,.32,1)
}

.flatpickr-calendar.inline {
    display: block;
    position: relative;
    top: 2px
}

.flatpickr-calendar.static {
    position: absolute;
    top: calc(100% + 2px)
}

.flatpickr-calendar.static.open {
    z-index: 999;
    display: block
}

.flatpickr-calendar.multiMonth .flatpickr-days .dayContainer:nth-child(n+1) .flatpickr-day.inRange:nth-child(7n+7) {
    box-shadow: none!important
}

.flatpickr-calendar.multiMonth .flatpickr-days .dayContainer:nth-child(n+2) .flatpickr-day.inRange:nth-child(7n+1) {
    box-shadow: -2px 0 #e6e6e6,5px 0 #e6e6e6
}

.flatpickr-calendar .hasWeeks .dayContainer,.flatpickr-calendar .hasTime .dayContainer {
    border-bottom: 0;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0
}

.flatpickr-calendar .hasWeeks .dayContainer {
    border-left: 0
}

.flatpickr-calendar.hasTime .flatpickr-time {
    height: 40px;
    border-top: 1px solid #e6e6e6
}

.flatpickr-calendar.noCalendar.hasTime .flatpickr-time {
    height: auto
}

.flatpickr-calendar:before,.flatpickr-calendar:after {
    position: absolute;
    display: block;
    pointer-events: none;
    border: solid transparent;
    content: "";
    height: 0;
    width: 0;
    left: 22px
}

.flatpickr-calendar.rightMost:before,.flatpickr-calendar.arrowRight:before,.flatpickr-calendar.rightMost:after,.flatpickr-calendar.arrowRight:after {
    left: auto;
    right: 22px
}

.flatpickr-calendar.arrowCenter:before,.flatpickr-calendar.arrowCenter:after {
    left: 50%;
    right: 50%
}

.flatpickr-calendar:before {
    border-width: 5px;
    margin: 0 -5px
}

.flatpickr-calendar:after {
    border-width: 4px;
    margin: 0 -4px
}

.flatpickr-calendar.arrowTop:before,.flatpickr-calendar.arrowTop:after {
    bottom: 100%
}

.flatpickr-calendar.arrowTop:before {
    border-bottom-color: #e6e6e6
}

.flatpickr-calendar.arrowTop:after {
    border-bottom-color: #fff
}

.flatpickr-calendar.arrowBottom:before,.flatpickr-calendar.arrowBottom:after {
    top: 100%
}

.flatpickr-calendar.arrowBottom:before {
    border-top-color: #e6e6e6
}

.flatpickr-calendar.arrowBottom:after {
    border-top-color: #fff
}

.flatpickr-calendar:focus {
    outline: 0
}

.flatpickr-wrapper {
    position: relative;
    display: inline-block
}

.flatpickr-months {
    display: flex
}

.flatpickr-months .flatpickr-month {
    background: transparent;
    color: #000000e6;
    fill: #000000e6;
    height: 34px;
    line-height: 1;
    text-align: center;
    position: relative;
    -webkit-user-select: none;
    user-select: none;
    overflow: hidden;
    flex: 1
}

.flatpickr-months .flatpickr-prev-month,.flatpickr-months .flatpickr-next-month {
    -webkit-user-select: none;
    user-select: none;
    text-decoration: none;
    cursor: pointer;
    position: absolute;
    top: 0;
    height: 34px;
    padding: 10px;
    z-index: 3;
    color: #000000e6;
    fill: #000000e6
}

.flatpickr-months .flatpickr-prev-month.flatpickr-disabled,.flatpickr-months .flatpickr-next-month.flatpickr-disabled {
    display: none
}

.flatpickr-months .flatpickr-prev-month i,.flatpickr-months .flatpickr-next-month i {
    position: relative
}

.flatpickr-months .flatpickr-prev-month.flatpickr-prev-month,.flatpickr-months .flatpickr-next-month.flatpickr-prev-month {
    left: 0
}

.flatpickr-months .flatpickr-prev-month.flatpickr-next-month,.flatpickr-months .flatpickr-next-month.flatpickr-next-month {
    right: 0
}

.flatpickr-months .flatpickr-prev-month:hover,.flatpickr-months .flatpickr-next-month:hover {
    color: #959ea9
}

.flatpickr-months .flatpickr-prev-month:hover svg,.flatpickr-months .flatpickr-next-month:hover svg {
    fill: #f64747
}

.flatpickr-months .flatpickr-prev-month svg,.flatpickr-months .flatpickr-next-month svg {
    width: 14px;
    height: 14px
}

.flatpickr-months .flatpickr-prev-month svg path,.flatpickr-months .flatpickr-next-month svg path {
    transition: fill .1s;
    fill: inherit
}

.numInputWrapper {
    position: relative;
    height: auto
}

.numInputWrapper input,.numInputWrapper span {
    display: inline-block
}

.numInputWrapper input {
    width: 100%
}

.numInputWrapper input::-ms-clear {
    display: none
}

.numInputWrapper input::-webkit-outer-spin-button,.numInputWrapper input::-webkit-inner-spin-button {
    margin: 0;
    -webkit-appearance: none
}

.numInputWrapper span {
    position: absolute;
    right: 0;
    width: 14px;
    padding: 0 4px 0 2px;
    height: 50%;
    line-height: 50%;
    opacity: 0;
    cursor: pointer;
    border: 1px solid rgba(57,57,57,.15);
    box-sizing: border-box
}

.numInputWrapper span:hover {
    background: rgba(0,0,0,.1)
}

.numInputWrapper span:active {
    background: rgba(0,0,0,.2)
}

.numInputWrapper span:after {
    display: block;
    content: "";
    position: absolute
}

.numInputWrapper span.arrowUp {
    top: 0;
    border-bottom: 0
}

.numInputWrapper span.arrowUp:after {
    border-left: 4px solid transparent;
    border-right: 4px solid transparent;
    border-bottom: 4px solid rgba(57,57,57,.6);
    top: 26%
}

.numInputWrapper span.arrowDown {
    top: 50%
}

.numInputWrapper span.arrowDown:after {
    border-left: 4px solid transparent;
    border-right: 4px solid transparent;
    border-top: 4px solid rgba(57,57,57,.6);
    top: 40%
}

.numInputWrapper span svg {
    width: inherit;
    height: auto
}

.numInputWrapper span svg path {
    fill: #00000080
}

.numInputWrapper:hover {
    background: rgba(0,0,0,.05)
}

.numInputWrapper:hover span {
    opacity: 1
}

.flatpickr-current-month {
    font-size: 135%;
    line-height: inherit;
    font-weight: 300;
    color: inherit;
    position: absolute;
    width: 75%;
    left: 12.5%;
    padding: 7.48px 0 0;
    line-height: 1;
    height: 34px;
    display: inline-block;
    text-align: center;
    transform: translateZ(0)
}

.flatpickr-current-month span.cur-month {
    font-family: inherit;
    font-weight: 700;
    color: inherit;
    display: inline-block;
    margin-left: .5ch;
    padding: 0
}

.flatpickr-current-month span.cur-month:hover {
    background: rgba(0,0,0,.05)
}

.flatpickr-current-month .numInputWrapper {
    width: 6ch;
    width: 7ch\fffd;
    display: inline-block
}

.flatpickr-current-month .numInputWrapper span.arrowUp:after {
    border-bottom-color: #000000e6
}

.flatpickr-current-month .numInputWrapper span.arrowDown:after {
    border-top-color: #000000e6
}

.flatpickr-current-month input.cur-year {
    background: transparent;
    box-sizing: border-box;
    color: inherit;
    cursor: text;
    padding: 0 0 0 .5ch;
    margin: 0;
    display: inline-block;
    font-size: inherit;
    font-family: inherit;
    font-weight: 300;
    line-height: inherit;
    height: auto;
    border: 0;
    border-radius: 0;
    vertical-align: initial;
    -webkit-appearance: textfield;
    appearance: textfield
}

.flatpickr-current-month input.cur-year:focus {
    outline: 0
}

.flatpickr-current-month input.cur-year[disabled],.flatpickr-current-month input.cur-year[disabled]:hover {
    font-size: 100%;
    color: #00000080;
    background: transparent;
    pointer-events: none
}

.flatpickr-current-month .flatpickr-monthDropdown-months {
    appearance: menulist;
    background: transparent;
    border: none;
    border-radius: 0;
    box-sizing: border-box;
    color: inherit;
    cursor: pointer;
    font-size: inherit;
    font-family: inherit;
    font-weight: 300;
    height: auto;
    line-height: inherit;
    margin: -1px 0 0;
    outline: none;
    padding: 0 0 0 .5ch;
    position: relative;
    vertical-align: initial;
    -webkit-box-sizing: border-box;
    -webkit-appearance: menulist;
    -moz-appearance: menulist;
    width: auto
}

.flatpickr-current-month .flatpickr-monthDropdown-months:focus,.flatpickr-current-month .flatpickr-monthDropdown-months:active {
    outline: none
}

.flatpickr-current-month .flatpickr-monthDropdown-months:hover {
    background: rgba(0,0,0,.05)
}

.flatpickr-current-month .flatpickr-monthDropdown-months .flatpickr-monthDropdown-month {
    background-color: transparent;
    outline: none;
    padding: 0
}

.flatpickr-weekdays {
    background: transparent;
    text-align: center;
    overflow: hidden;
    width: 100%;
    display: flex;
    align-items: center;
    height: 28px
}

.flatpickr-weekdays .flatpickr-weekdaycontainer {
    display: flex;
    flex: 1
}

span.flatpickr-weekday {
    cursor: default;
    font-size: 90%;
    background: transparent;
    color: #0000008a;
    line-height: 1;
    margin: 0;
    text-align: center;
    display: block;
    flex: 1;
    font-weight: bolder
}

.dayContainer,.flatpickr-weeks {
    padding: 1px 0 0
}

.flatpickr-days {
    position: relative;
    overflow: hidden;
    display: flex;
    align-items: flex-start;
    width: 307.875px
}

.flatpickr-days:focus {
    outline: 0
}

.dayContainer {
    padding: 0;
    outline: 0;
    text-align: left;
    width: 307.875px;
    min-width: 307.875px;
    max-width: 307.875px;
    box-sizing: border-box;
    display: inline-block;
    display: flex;
    flex-wrap: wrap;
    -ms-flex-wrap: wrap;
    justify-content: space-around;
    transform: translateZ(0);
    opacity: 1
}

.dayContainer+.dayContainer {
    box-shadow: -1px 0 #e6e6e6
}

.flatpickr-day {
    background: none;
    border: 1px solid transparent;
    border-radius: 150px;
    box-sizing: border-box;
    color: #393939;
    cursor: pointer;
    font-weight: 400;
    width: 14.2857143%;
    flex-basis: 14.2857143%;
    max-width: 39px;
    height: 39px;
    line-height: 39px;
    margin: 0;
    display: inline-block;
    position: relative;
    justify-content: center;
    text-align: center
}

.flatpickr-day.inRange,.flatpickr-day.prevMonthDay.inRange,.flatpickr-day.nextMonthDay.inRange,.flatpickr-day.today.inRange,.flatpickr-day.prevMonthDay.today.inRange,.flatpickr-day.nextMonthDay.today.inRange,.flatpickr-day:hover,.flatpickr-day.prevMonthDay:hover,.flatpickr-day.nextMonthDay:hover,.flatpickr-day:focus,.flatpickr-day.prevMonthDay:focus,.flatpickr-day.nextMonthDay:focus {
    cursor: pointer;
    outline: 0;
    background: #e6e6e6;
    border-color: #e6e6e6
}

.flatpickr-day.today {
    border-color: #959ea9
}

.flatpickr-day.today:hover,.flatpickr-day.today:focus {
    border-color: #959ea9;
    background: #959ea9;
    color: #fff
}

.flatpickr-day.selected,.flatpickr-day.startRange,.flatpickr-day.endRange,.flatpickr-day.selected.inRange,.flatpickr-day.startRange.inRange,.flatpickr-day.endRange.inRange,.flatpickr-day.selected:focus,.flatpickr-day.startRange:focus,.flatpickr-day.endRange:focus,.flatpickr-day.selected:hover,.flatpickr-day.startRange:hover,.flatpickr-day.endRange:hover,.flatpickr-day.selected.prevMonthDay,.flatpickr-day.startRange.prevMonthDay,.flatpickr-day.endRange.prevMonthDay,.flatpickr-day.selected.nextMonthDay,.flatpickr-day.startRange.nextMonthDay,.flatpickr-day.endRange.nextMonthDay {
    background: #569ff7;
    box-shadow: none;
    color: #fff;
    border-color: #569ff7
}

.flatpickr-day.selected.startRange,.flatpickr-day.startRange.startRange,.flatpickr-day.endRange.startRange {
    border-radius: 50px 0 0 50px
}

.flatpickr-day.selected.endRange,.flatpickr-day.startRange.endRange,.flatpickr-day.endRange.endRange {
    border-radius: 0 50px 50px 0
}

.flatpickr-day.selected.startRange+.endRange:not(:nth-child(7n+1)),.flatpickr-day.startRange.startRange+.endRange:not(:nth-child(7n+1)),.flatpickr-day.endRange.startRange+.endRange:not(:nth-child(7n+1)) {
    box-shadow: -10px 0 #569ff7
}

.flatpickr-day.selected.startRange.endRange,.flatpickr-day.startRange.startRange.endRange,.flatpickr-day.endRange.startRange.endRange {
    border-radius: 50px
}

.flatpickr-day.inRange {
    border-radius: 0;
    box-shadow: -5px 0 #e6e6e6,5px 0 #e6e6e6
}

.flatpickr-day.flatpickr-disabled,.flatpickr-day.flatpickr-disabled:hover,.flatpickr-day.prevMonthDay,.flatpickr-day.nextMonthDay,.flatpickr-day.notAllowed,.flatpickr-day.notAllowed.prevMonthDay,.flatpickr-day.notAllowed.nextMonthDay {
    color: #3939394d;
    background: transparent;
    border-color: transparent;
    cursor: default
}

.flatpickr-day.flatpickr-disabled,.flatpickr-day.flatpickr-disabled:hover {
    cursor: not-allowed;
    color: #3939391a
}

.flatpickr-day.week.selected {
    border-radius: 0;
    box-shadow: -5px 0 #569ff7,5px 0 #569ff7
}

.flatpickr-day.hidden {
    visibility: hidden
}

.rangeMode .flatpickr-day {
    margin-top: 1px
}

.flatpickr-weekwrapper {
    float: left
}

.flatpickr-weekwrapper .flatpickr-weeks {
    padding: 0 12px;
    box-shadow: 1px 0 #e6e6e6
}

.flatpickr-weekwrapper .flatpickr-weekday {
    float: none;
    width: 100%;
    line-height: 28px
}

.flatpickr-weekwrapper span.flatpickr-day,.flatpickr-weekwrapper span.flatpickr-day:hover {
    display: block;
    width: 100%;
    max-width: none;
    color: #3939394d;
    background: transparent;
    cursor: default;
    border: none
}

.flatpickr-innerContainer {
    display: block;
    display: flex;
    box-sizing: border-box;
    overflow: hidden
}

.flatpickr-rContainer {
    display: inline-block;
    padding: 0;
    box-sizing: border-box
}

.flatpickr-time {
    text-align: center;
    outline: 0;
    display: block;
    height: 0;
    line-height: 40px;
    max-height: 40px;
    box-sizing: border-box;
    overflow: hidden;
    display: flex
}

.flatpickr-time:after {
    content: "";
    display: table;
    clear: both
}

.flatpickr-time .numInputWrapper {
    flex: 1;
    width: 40%;
    height: 40px;
    float: left
}

.flatpickr-time .numInputWrapper span.arrowUp:after {
    border-bottom-color: #393939
}

.flatpickr-time .numInputWrapper span.arrowDown:after {
    border-top-color: #393939
}

.flatpickr-time.hasSeconds .numInputWrapper {
    width: 26%
}

.flatpickr-time.time24hr .numInputWrapper {
    width: 49%
}

.flatpickr-time input {
    background: transparent;
    box-shadow: none;
    border: 0;
    border-radius: 0;
    text-align: center;
    margin: 0;
    padding: 0;
    height: inherit;
    line-height: inherit;
    color: #393939;
    font-size: 14px;
    position: relative;
    box-sizing: border-box;
    -webkit-appearance: textfield;
    appearance: textfield
}

.flatpickr-time input.flatpickr-hour {
    font-weight: 700
}

.flatpickr-time input.flatpickr-minute,.flatpickr-time input.flatpickr-second {
    font-weight: 400
}

.flatpickr-time input:focus {
    outline: 0;
    border: 0
}

.flatpickr-time .flatpickr-time-separator,.flatpickr-time .flatpickr-am-pm {
    height: inherit;
    float: left;
    line-height: inherit;
    color: #393939;
    font-weight: 700;
    width: 2%;
    -webkit-user-select: none;
    user-select: none;
    align-self: center
}

.flatpickr-time .flatpickr-am-pm {
    outline: 0;
    width: 18%;
    cursor: pointer;
    text-align: center;
    font-weight: 400
}

.flatpickr-time input:hover,.flatpickr-time .flatpickr-am-pm:hover,.flatpickr-time input:focus,.flatpickr-time .flatpickr-am-pm:focus {
    background: #eee
}

.flatpickr-input[readonly] {
    cursor: pointer
}

@keyframes fpFadeInDown {
    0% {
        opacity: 0;
        transform: translate3d(0,-20px,0)
    }

    to {
        opacity: 1;
        transform: translateZ(0)
    }
}

/*!
 * Quill Editor v2.0.18
 * https://quilljs.com/
 * Copyright (c) 2014, Jason Chen
 * Copyright (c) 2013, salesforce.com
 */
.ql-container {
    box-sizing: border-box;
    font-family: Helvetica,Arial,sans-serif;
    font-size: 13px;
    height: 100%;
    margin: 0;
    position: relative
}

.ql-container.ql-disabled .ql-tooltip {
    visibility: hidden
}

.ql-container:not(.ql-disabled) li[data-list=checked]>.ql-ui,.ql-container:not(.ql-disabled) li[data-list=unchecked]>.ql-ui {
    cursor: pointer
}

.ql-clipboard {
    left: -100000px;
    height: 1px;
    overflow-y: hidden;
    position: absolute;
    top: 50%
}

.ql-clipboard p {
    margin: 0;
    padding: 0
}

.ql-editor {
    box-sizing: border-box;
    counter-reset: list-0 list-1 list-2 list-3 list-4 list-5 list-6 list-7 list-8 list-9;
    line-height: 1.42;
    height: 100%;
    outline: none;
    overflow-y: auto;
    padding: 12px 15px;
    tab-size: 4;
    -moz-tab-size: 4;
    text-align: left;
    white-space: pre-wrap;
    word-wrap: break-word
}

.ql-editor>* {
    cursor: text
}

.ql-editor p,.ql-editor ol,.ql-editor pre,.ql-editor blockquote,.ql-editor h1,.ql-editor h2,.ql-editor h3,.ql-editor h4,.ql-editor h5,.ql-editor h6 {
    margin: 0;
    padding: 0
}

@supports (counter-set: none) {
    .ql-editor p,.ql-editor h1,.ql-editor h2,.ql-editor h3,.ql-editor h4,.ql-editor h5,.ql-editor h6 {
        counter-set: list-0 list-1 list-2 list-3 list-4 list-5 list-6 list-7 list-8 list-9
    }
}

@supports not (counter-set: none) {
    .ql-editor p,.ql-editor h1,.ql-editor h2,.ql-editor h3,.ql-editor h4,.ql-editor h5,.ql-editor h6 {
        counter-reset: list-0 list-1 list-2 list-3 list-4 list-5 list-6 list-7 list-8 list-9
    }
}

.ql-editor table {
    border-collapse: collapse
}

.ql-editor td {
    border: 1px solid #000;
    padding: 2px 5px
}

.ql-editor ol {
    padding-left: 1.5em
}

.ql-editor li {
    list-style-type: none;
    padding-left: 1.5em;
    position: relative
}

.ql-editor li>.ql-ui:before {
    display: inline-block;
    margin-left: -1.5em;
    margin-right: .3em;
    text-align: right;
    white-space: nowrap;
    width: 1.2em
}

@supports (display: contents) {
    .ql-editor li[data-list=bullet]>.ql-ui,.ql-editor li[data-list=ordered]>.ql-ui {
        display:contents
    }
}

.ql-editor li[data-list=checked]>.ql-ui,.ql-editor li[data-list=unchecked]>.ql-ui {
    color: #777
}

.ql-editor li[data-list=bullet]>.ql-ui:before {
    content: "\2022"
}

.ql-editor li[data-list=checked]>.ql-ui:before {
    content: "\2611"
}

.ql-editor li[data-list=unchecked]>.ql-ui:before {
    content: "\2610"
}

@supports (counter-set: none) {
    .ql-editor li[data-list] {
        counter-set: list-1 list-2 list-3 list-4 list-5 list-6 list-7 list-8 list-9
    }
}

@supports not (counter-set: none) {
    .ql-editor li[data-list] {
        counter-reset: list-1 list-2 list-3 list-4 list-5 list-6 list-7 list-8 list-9
    }
}

.ql-editor li[data-list=ordered] {
    counter-increment: list-0
}

.ql-editor li[data-list=ordered]>.ql-ui:before {
    content: counter(list-0,decimal) ". "
}

.ql-editor li[data-list=ordered].ql-indent-1 {
    counter-increment: list-1
}

.ql-editor li[data-list=ordered].ql-indent-1>.ql-ui:before {
    content: counter(list-1,lower-alpha) ". "
}

@supports (counter-set: none) {
    .ql-editor li[data-list].ql-indent-1 {
        counter-set: list-2 list-3 list-4 list-5 list-6 list-7 list-8 list-9
    }
}

@supports not (counter-set: none) {
    .ql-editor li[data-list].ql-indent-1 {
        counter-reset: list-2 list-3 list-4 list-5 list-6 list-7 list-8 list-9
    }
}

.ql-editor li[data-list=ordered].ql-indent-2 {
    counter-increment: list-2
}

.ql-editor li[data-list=ordered].ql-indent-2>.ql-ui:before {
    content: counter(list-2,lower-roman) ". "
}

@supports (counter-set: none) {
    .ql-editor li[data-list].ql-indent-2 {
        counter-set: list-3 list-4 list-5 list-6 list-7 list-8 list-9
    }
}

@supports not (counter-set: none) {
    .ql-editor li[data-list].ql-indent-2 {
        counter-reset: list-3 list-4 list-5 list-6 list-7 list-8 list-9
    }
}

.ql-editor li[data-list=ordered].ql-indent-3 {
    counter-increment: list-3
}

.ql-editor li[data-list=ordered].ql-indent-3>.ql-ui:before {
    content: counter(list-3,decimal) ". "
}

@supports (counter-set: none) {
    .ql-editor li[data-list].ql-indent-3 {
        counter-set: list-4 list-5 list-6 list-7 list-8 list-9
    }
}

@supports not (counter-set: none) {
    .ql-editor li[data-list].ql-indent-3 {
        counter-reset: list-4 list-5 list-6 list-7 list-8 list-9
    }
}

.ql-editor li[data-list=ordered].ql-indent-4 {
    counter-increment: list-4
}

.ql-editor li[data-list=ordered].ql-indent-4>.ql-ui:before {
    content: counter(list-4,lower-alpha) ". "
}

@supports (counter-set: none) {
    .ql-editor li[data-list].ql-indent-4 {
        counter-set: list-5 list-6 list-7 list-8 list-9
    }
}

@supports not (counter-set: none) {
    .ql-editor li[data-list].ql-indent-4 {
        counter-reset: list-5 list-6 list-7 list-8 list-9
    }
}

.ql-editor li[data-list=ordered].ql-indent-5 {
    counter-increment: list-5
}

.ql-editor li[data-list=ordered].ql-indent-5>.ql-ui:before {
    content: counter(list-5,lower-roman) ". "
}

@supports (counter-set: none) {
    .ql-editor li[data-list].ql-indent-5 {
        counter-set: list-6 list-7 list-8 list-9
    }
}

@supports not (counter-set: none) {
    .ql-editor li[data-list].ql-indent-5 {
        counter-reset: list-6 list-7 list-8 list-9
    }
}

.ql-editor li[data-list=ordered].ql-indent-6 {
    counter-increment: list-6
}

.ql-editor li[data-list=ordered].ql-indent-6>.ql-ui:before {
    content: counter(list-6,decimal) ". "
}

@supports (counter-set: none) {
    .ql-editor li[data-list].ql-indent-6 {
        counter-set: list-7 list-8 list-9
    }
}

@supports not (counter-set: none) {
    .ql-editor li[data-list].ql-indent-6 {
        counter-reset: list-7 list-8 list-9
    }
}

.ql-editor li[data-list=ordered].ql-indent-7 {
    counter-increment: list-7
}

.ql-editor li[data-list=ordered].ql-indent-7>.ql-ui:before {
    content: counter(list-7,lower-alpha) ". "
}

@supports (counter-set: none) {
    .ql-editor li[data-list].ql-indent-7 {
        counter-set: list-8 list-9
    }
}

@supports not (counter-set: none) {
    .ql-editor li[data-list].ql-indent-7 {
        counter-reset: list-8 list-9
    }
}

.ql-editor li[data-list=ordered].ql-indent-8 {
    counter-increment: list-8
}

.ql-editor li[data-list=ordered].ql-indent-8>.ql-ui:before {
    content: counter(list-8,lower-roman) ". "
}

@supports (counter-set: none) {
    .ql-editor li[data-list].ql-indent-8 {
        counter-set: list-9
    }
}

@supports not (counter-set: none) {
    .ql-editor li[data-list].ql-indent-8 {
        counter-reset: list-9
    }
}

.ql-editor li[data-list=ordered].ql-indent-9 {
    counter-increment: list-9
}

.ql-editor li[data-list=ordered].ql-indent-9>.ql-ui:before {
    content: counter(list-9,decimal) ". "
}

.ql-editor .ql-indent-1:not(.ql-direction-rtl) {
    padding-left: 3em
}

.ql-editor li.ql-indent-1:not(.ql-direction-rtl) {
    padding-left: 4.5em
}

.ql-editor .ql-indent-1.ql-direction-rtl.ql-align-right {
    padding-right: 3em
}

.ql-editor li.ql-indent-1.ql-direction-rtl.ql-align-right {
    padding-right: 4.5em
}

.ql-editor .ql-indent-2:not(.ql-direction-rtl) {
    padding-left: 6em
}

.ql-editor li.ql-indent-2:not(.ql-direction-rtl) {
    padding-left: 7.5em
}

.ql-editor .ql-indent-2.ql-direction-rtl.ql-align-right {
    padding-right: 6em
}

.ql-editor li.ql-indent-2.ql-direction-rtl.ql-align-right {
    padding-right: 7.5em
}

.ql-editor .ql-indent-3:not(.ql-direction-rtl) {
    padding-left: 9em
}

.ql-editor li.ql-indent-3:not(.ql-direction-rtl) {
    padding-left: 10.5em
}

.ql-editor .ql-indent-3.ql-direction-rtl.ql-align-right {
    padding-right: 9em
}

.ql-editor li.ql-indent-3.ql-direction-rtl.ql-align-right {
    padding-right: 10.5em
}

.ql-editor .ql-indent-4:not(.ql-direction-rtl) {
    padding-left: 12em
}

.ql-editor li.ql-indent-4:not(.ql-direction-rtl) {
    padding-left: 13.5em
}

.ql-editor .ql-indent-4.ql-direction-rtl.ql-align-right {
    padding-right: 12em
}

.ql-editor li.ql-indent-4.ql-direction-rtl.ql-align-right {
    padding-right: 13.5em
}

.ql-editor .ql-indent-5:not(.ql-direction-rtl) {
    padding-left: 15em
}

.ql-editor li.ql-indent-5:not(.ql-direction-rtl) {
    padding-left: 16.5em
}

.ql-editor .ql-indent-5.ql-direction-rtl.ql-align-right {
    padding-right: 15em
}

.ql-editor li.ql-indent-5.ql-direction-rtl.ql-align-right {
    padding-right: 16.5em
}

.ql-editor .ql-indent-6:not(.ql-direction-rtl) {
    padding-left: 18em
}

.ql-editor li.ql-indent-6:not(.ql-direction-rtl) {
    padding-left: 19.5em
}

.ql-editor .ql-indent-6.ql-direction-rtl.ql-align-right {
    padding-right: 18em
}

.ql-editor li.ql-indent-6.ql-direction-rtl.ql-align-right {
    padding-right: 19.5em
}

.ql-editor .ql-indent-7:not(.ql-direction-rtl) {
    padding-left: 21em
}

.ql-editor li.ql-indent-7:not(.ql-direction-rtl) {
    padding-left: 22.5em
}

.ql-editor .ql-indent-7.ql-direction-rtl.ql-align-right {
    padding-right: 21em
}

.ql-editor li.ql-indent-7.ql-direction-rtl.ql-align-right {
    padding-right: 22.5em
}

.ql-editor .ql-indent-8:not(.ql-direction-rtl) {
    padding-left: 24em
}

.ql-editor li.ql-indent-8:not(.ql-direction-rtl) {
    padding-left: 25.5em
}

.ql-editor .ql-indent-8.ql-direction-rtl.ql-align-right {
    padding-right: 24em
}

.ql-editor li.ql-indent-8.ql-direction-rtl.ql-align-right {
    padding-right: 25.5em
}

.ql-editor .ql-indent-9:not(.ql-direction-rtl) {
    padding-left: 27em
}

.ql-editor li.ql-indent-9:not(.ql-direction-rtl) {
    padding-left: 28.5em
}

.ql-editor .ql-indent-9.ql-direction-rtl.ql-align-right {
    padding-right: 27em
}

.ql-editor li.ql-indent-9.ql-direction-rtl.ql-align-right {
    padding-right: 28.5em
}

.ql-editor li.ql-direction-rtl {
    padding-right: 1.5em
}

.ql-editor li.ql-direction-rtl>.ql-ui:before {
    margin-left: .3em;
    margin-right: -1.5em;
    text-align: left
}

.ql-editor table {
    table-layout: fixed;
    width: 100%
}

.ql-editor table td {
    outline: none
}

.ql-editor .ql-code-block-container {
    font-family: monospace
}

.ql-editor .ql-video {
    display: block;
    max-width: 100%
}

.ql-editor .ql-video.ql-align-center {
    margin: 0 auto
}

.ql-editor .ql-video.ql-align-right {
    margin: 0 0 0 auto
}

.ql-editor .ql-bg-black {
    background-color: #000
}

.ql-editor .ql-bg-red {
    background-color: #e60000
}

.ql-editor .ql-bg-orange {
    background-color: #f90
}

.ql-editor .ql-bg-yellow {
    background-color: #ff0
}

.ql-editor .ql-bg-green {
    background-color: #008a00
}

.ql-editor .ql-bg-blue {
    background-color: #06c
}

.ql-editor .ql-bg-purple {
    background-color: #93f
}

.ql-editor .ql-color-white {
    color: #fff
}

.ql-editor .ql-color-red {
    color: #e60000
}

.ql-editor .ql-color-orange {
    color: #f90
}

.ql-editor .ql-color-yellow {
    color: #ff0
}

.ql-editor .ql-color-green {
    color: #008a00
}

.ql-editor .ql-color-blue {
    color: #06c
}

.ql-editor .ql-color-purple {
    color: #93f
}

.ql-editor .ql-font-serif {
    font-family: Georgia,Times New Roman,serif
}

.ql-editor .ql-font-monospace {
    font-family: Monaco,Courier New,monospace
}

.ql-editor .ql-size-small {
    font-size: .75em
}

.ql-editor .ql-size-large {
    font-size: 1.5em
}

.ql-editor .ql-size-huge {
    font-size: 2.5em
}

.ql-editor .ql-direction-rtl {
    direction: rtl;
    text-align: inherit
}

.ql-editor .ql-align-center {
    text-align: center
}

.ql-editor .ql-align-justify {
    text-align: justify
}

.ql-editor .ql-align-right {
    text-align: right
}

.ql-editor .ql-ui {
    position: absolute
}

.ql-editor.ql-blank:before {
    color: #0009;
    content: attr(data-placeholder);
    font-style: italic;
    left: 15px;
    pointer-events: none;
    position: absolute;
    right: 15px
}

.ql-snow.ql-toolbar:after,.ql-snow .ql-toolbar:after {
    clear: both;
    content: "";
    display: table
}

.ql-snow.ql-toolbar button,.ql-snow .ql-toolbar button {
    background: none;
    border: none;
    cursor: pointer;
    display: inline-block;
    float: left;
    height: 24px;
    padding: 3px 5px;
    width: 28px
}

.ql-snow.ql-toolbar button svg,.ql-snow .ql-toolbar button svg {
    float: left;
    height: 100%
}

.ql-snow.ql-toolbar button:active:hover,.ql-snow .ql-toolbar button:active:hover {
    outline: none
}

.ql-snow.ql-toolbar input.ql-image[type=file],.ql-snow .ql-toolbar input.ql-image[type=file] {
    display: none
}

.ql-snow.ql-toolbar button:hover,.ql-snow .ql-toolbar button:hover,.ql-snow.ql-toolbar button:focus,.ql-snow .ql-toolbar button:focus,.ql-snow.ql-toolbar button.ql-active,.ql-snow .ql-toolbar button.ql-active,.ql-snow.ql-toolbar .ql-picker-label:hover,.ql-snow .ql-toolbar .ql-picker-label:hover,.ql-snow.ql-toolbar .ql-picker-label.ql-active,.ql-snow .ql-toolbar .ql-picker-label.ql-active,.ql-snow.ql-toolbar .ql-picker-item:hover,.ql-snow .ql-toolbar .ql-picker-item:hover,.ql-snow.ql-toolbar .ql-picker-item.ql-selected,.ql-snow .ql-toolbar .ql-picker-item.ql-selected {
    color: #06c
}

.ql-snow.ql-toolbar button:hover .ql-fill,.ql-snow .ql-toolbar button:hover .ql-fill,.ql-snow.ql-toolbar button:focus .ql-fill,.ql-snow .ql-toolbar button:focus .ql-fill,.ql-snow.ql-toolbar button.ql-active .ql-fill,.ql-snow .ql-toolbar button.ql-active .ql-fill,.ql-snow.ql-toolbar .ql-picker-label:hover .ql-fill,.ql-snow .ql-toolbar .ql-picker-label:hover .ql-fill,.ql-snow.ql-toolbar .ql-picker-label.ql-active .ql-fill,.ql-snow .ql-toolbar .ql-picker-label.ql-active .ql-fill,.ql-snow.ql-toolbar .ql-picker-item:hover .ql-fill,.ql-snow .ql-toolbar .ql-picker-item:hover .ql-fill,.ql-snow.ql-toolbar .ql-picker-item.ql-selected .ql-fill,.ql-snow .ql-toolbar .ql-picker-item.ql-selected .ql-fill,.ql-snow.ql-toolbar button:hover .ql-stroke.ql-fill,.ql-snow .ql-toolbar button:hover .ql-stroke.ql-fill,.ql-snow.ql-toolbar button:focus .ql-stroke.ql-fill,.ql-snow .ql-toolbar button:focus .ql-stroke.ql-fill,.ql-snow.ql-toolbar button.ql-active .ql-stroke.ql-fill,.ql-snow .ql-toolbar button.ql-active .ql-stroke.ql-fill,.ql-snow.ql-toolbar .ql-picker-label:hover .ql-stroke.ql-fill,.ql-snow .ql-toolbar .ql-picker-label:hover .ql-stroke.ql-fill,.ql-snow.ql-toolbar .ql-picker-label.ql-active .ql-stroke.ql-fill,.ql-snow .ql-toolbar .ql-picker-label.ql-active .ql-stroke.ql-fill,.ql-snow.ql-toolbar .ql-picker-item:hover .ql-stroke.ql-fill,.ql-snow .ql-toolbar .ql-picker-item:hover .ql-stroke.ql-fill,.ql-snow.ql-toolbar .ql-picker-item.ql-selected .ql-stroke.ql-fill,.ql-snow .ql-toolbar .ql-picker-item.ql-selected .ql-stroke.ql-fill {
    fill: #06c
}

.ql-snow.ql-toolbar button:hover .ql-stroke,.ql-snow .ql-toolbar button:hover .ql-stroke,.ql-snow.ql-toolbar button:focus .ql-stroke,.ql-snow .ql-toolbar button:focus .ql-stroke,.ql-snow.ql-toolbar button.ql-active .ql-stroke,.ql-snow .ql-toolbar button.ql-active .ql-stroke,.ql-snow.ql-toolbar .ql-picker-label:hover .ql-stroke,.ql-snow .ql-toolbar .ql-picker-label:hover .ql-stroke,.ql-snow.ql-toolbar .ql-picker-label.ql-active .ql-stroke,.ql-snow .ql-toolbar .ql-picker-label.ql-active .ql-stroke,.ql-snow.ql-toolbar .ql-picker-item:hover .ql-stroke,.ql-snow .ql-toolbar .ql-picker-item:hover .ql-stroke,.ql-snow.ql-toolbar .ql-picker-item.ql-selected .ql-stroke,.ql-snow .ql-toolbar .ql-picker-item.ql-selected .ql-stroke,.ql-snow.ql-toolbar button:hover .ql-stroke-miter,.ql-snow .ql-toolbar button:hover .ql-stroke-miter,.ql-snow.ql-toolbar button:focus .ql-stroke-miter,.ql-snow .ql-toolbar button:focus .ql-stroke-miter,.ql-snow.ql-toolbar button.ql-active .ql-stroke-miter,.ql-snow .ql-toolbar button.ql-active .ql-stroke-miter,.ql-snow.ql-toolbar .ql-picker-label:hover .ql-stroke-miter,.ql-snow .ql-toolbar .ql-picker-label:hover .ql-stroke-miter,.ql-snow.ql-toolbar .ql-picker-label.ql-active .ql-stroke-miter,.ql-snow .ql-toolbar .ql-picker-label.ql-active .ql-stroke-miter,.ql-snow.ql-toolbar .ql-picker-item:hover .ql-stroke-miter,.ql-snow .ql-toolbar .ql-picker-item:hover .ql-stroke-miter,.ql-snow.ql-toolbar .ql-picker-item.ql-selected .ql-stroke-miter,.ql-snow .ql-toolbar .ql-picker-item.ql-selected .ql-stroke-miter {
    stroke: #06c
}

@media (pointer: coarse) {
    .ql-snow.ql-toolbar button:hover:not(.ql-active),.ql-snow .ql-toolbar button:hover:not(.ql-active) {
        color:#444
    }

    .ql-snow.ql-toolbar button:hover:not(.ql-active) .ql-fill,.ql-snow .ql-toolbar button:hover:not(.ql-active) .ql-fill,.ql-snow.ql-toolbar button:hover:not(.ql-active) .ql-stroke.ql-fill,.ql-snow .ql-toolbar button:hover:not(.ql-active) .ql-stroke.ql-fill {
        fill: #444
    }

    .ql-snow.ql-toolbar button:hover:not(.ql-active) .ql-stroke,.ql-snow .ql-toolbar button:hover:not(.ql-active) .ql-stroke,.ql-snow.ql-toolbar button:hover:not(.ql-active) .ql-stroke-miter,.ql-snow .ql-toolbar button:hover:not(.ql-active) .ql-stroke-miter {
        stroke: #444
    }
}

.ql-snow,.ql-snow * {
    box-sizing: border-box
}

.ql-snow .ql-hidden {
    display: none
}

.ql-snow .ql-out-bottom,.ql-snow .ql-out-top {
    visibility: hidden
}

.ql-snow .ql-tooltip {
    position: absolute;
    transform: translateY(10px)
}

.ql-snow .ql-tooltip a {
    cursor: pointer;
    text-decoration: none
}

.ql-snow .ql-tooltip.ql-flip {
    transform: translateY(-10px)
}

.ql-snow .ql-formats {
    display: inline-block;
    vertical-align: middle
}

.ql-snow .ql-formats:after {
    clear: both;
    content: "";
    display: table
}

.ql-snow .ql-stroke {
    fill: none;
    stroke: #444;
    stroke-linecap: round;
    stroke-linejoin: round;
    stroke-width: 2
}

.ql-snow .ql-stroke-miter {
    fill: none;
    stroke: #444;
    stroke-miterlimit: 10;
    stroke-width: 2
}

.ql-snow .ql-fill,.ql-snow .ql-stroke.ql-fill {
    fill: #444
}

.ql-snow .ql-empty {
    fill: none
}

.ql-snow .ql-even {
    fill-rule: evenodd
}

.ql-snow .ql-thin,.ql-snow .ql-stroke.ql-thin {
    stroke-width: 1
}

.ql-snow .ql-transparent {
    opacity: .4
}

.ql-snow .ql-direction svg:last-child {
    display: none
}

.ql-snow .ql-direction.ql-active svg:last-child {
    display: inline
}

.ql-snow .ql-direction.ql-active svg:first-child {
    display: none
}

.ql-snow .ql-editor h1 {
    font-size: 2em
}

.ql-snow .ql-editor h2 {
    font-size: 1.5em
}

.ql-snow .ql-editor h3 {
    font-size: 1.17em
}

.ql-snow .ql-editor h4 {
    font-size: 1em
}

.ql-snow .ql-editor h5 {
    font-size: .83em
}

.ql-snow .ql-editor h6 {
    font-size: .67em
}

.ql-snow .ql-editor a {
    text-decoration: underline
}

.ql-snow .ql-editor blockquote {
    border-left: 4px solid #ccc;
    margin-bottom: 5px;
    margin-top: 5px;
    padding-left: 16px
}

.ql-snow .ql-editor code,.ql-snow .ql-editor .ql-code-block-container {
    background-color: #f0f0f0;
    border-radius: 3px
}

.ql-snow .ql-editor .ql-code-block-container {
    margin-bottom: 5px;
    margin-top: 5px;
    padding: 5px 10px
}

.ql-snow .ql-editor code {
    font-size: 85%;
    padding: 2px 4px
}

.ql-snow .ql-editor .ql-code-block-container {
    background-color: #23241f;
    color: #f8f8f2;
    overflow: visible
}

.ql-snow .ql-editor img {
    max-width: 100%
}

.ql-snow .ql-editor .ql-line-height-1 {
    line-height: 1
}

.ql-snow .ql-editor .ql-line-height-1_5 {
    line-height: 1.5
}

.ql-snow .ql-editor .ql-line-height-1_6 {
    line-height: 1.6
}

.ql-snow .ql-editor .ql-line-height-1_75 {
    line-height: 1.75
}

.ql-snow .ql-editor .ql-line-height-2 {
    line-height: 2
}

.ql-snow .ql-editor .ql-line-height-2_5 {
    line-height: 2.5
}

.ql-snow .ql-editor .ql-line-height-3 {
    line-height: 3
}

.ql-snow .ql-editor .ql-line-height-4 {
    line-height: 4
}

.ql-snow .ql-picker {
    color: #444;
    display: inline-block;
    float: left;
    font-size: 14px;
    font-weight: 500;
    height: 24px;
    position: relative;
    vertical-align: middle
}

.ql-snow .ql-picker-label {
    cursor: pointer;
    display: inline-block;
    height: 100%;
    padding-left: 8px;
    padding-right: 2px;
    position: relative;
    width: 100%
}

.ql-snow .ql-picker-label:before {
    display: inline-block;
    line-height: 22px
}

.ql-snow .ql-picker-options {
    background-color: #fff;
    display: none;
    min-width: 100%;
    padding: 4px 8px;
    position: absolute;
    white-space: nowrap
}

.ql-snow .ql-picker-options .ql-picker-item {
    cursor: pointer;
    display: block;
    padding-bottom: 5px;
    padding-top: 5px
}

.ql-snow .ql-picker.ql-expanded .ql-picker-label {
    color: #ccc;
    z-index: 2
}

.ql-snow .ql-picker.ql-expanded .ql-picker-label .ql-fill {
    fill: #ccc
}

.ql-snow .ql-picker.ql-expanded .ql-picker-label .ql-stroke {
    stroke: #ccc
}

.ql-snow .ql-picker.ql-expanded .ql-picker-options {
    display: block;
    margin-top: -1px;
    top: 100%;
    z-index: 1
}

.ql-snow .ql-color-picker,.ql-snow .ql-icon-picker,.ql-snow .ql-line-height {
    width: 28px
}

.ql-snow .ql-color-picker .ql-picker-label,.ql-snow .ql-icon-picker .ql-picker-label,.ql-snow .ql-line-height .ql-picker-label {
    padding: 2px 4px
}

.ql-snow .ql-color-picker .ql-picker-label svg,.ql-snow .ql-icon-picker .ql-picker-label svg,.ql-snow .ql-line-height .ql-picker-label svg {
    right: 4px
}

.ql-snow .ql-icon-picker .ql-picker-options {
    padding: 4px 0
}

.ql-snow .ql-icon-picker .ql-picker-item {
    height: 24px;
    width: 24px;
    padding: 2px 4px
}

.ql-snow .ql-color-picker .ql-picker-options {
    padding: 3px 5px;
    width: 152px
}

.ql-snow .ql-color-picker .ql-picker-item {
    border: 1px solid transparent;
    float: left;
    height: 16px;
    margin: 2px;
    padding: 0;
    width: 16px
}

.ql-snow .ql-picker:not(.ql-color-picker):not(.ql-icon-picker):not(.ql-line-height) svg {
    position: absolute;
    margin-top: -9px;
    right: 0;
    top: 50%;
    width: 18px
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-label]:not([data-label=""]):before,.ql-snow .ql-picker.ql-font .ql-picker-label[data-label]:not([data-label=""]):before,.ql-snow .ql-picker.ql-size .ql-picker-label[data-label]:not([data-label=""]):before,.ql-snow .ql-picker.ql-header .ql-picker-item[data-label]:not([data-label=""]):before,.ql-snow .ql-picker.ql-font .ql-picker-item[data-label]:not([data-label=""]):before,.ql-snow .ql-picker.ql-size .ql-picker-item[data-label]:not([data-label=""]):before {
    content: attr(data-label)
}

.ql-snow .ql-picker.ql-header {
    width: 98px
}

.ql-snow .ql-picker.ql-header .ql-picker-label:before,.ql-snow .ql-picker.ql-header .ql-picker-item:before {
    content: "Normal"
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="1"]:before,.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="1"]:before {
    content: "Heading 1"
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="2"]:before,.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="2"]:before {
    content: "Heading 2"
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="3"]:before,.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="3"]:before {
    content: "Heading 3"
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="4"]:before,.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="4"]:before {
    content: "Heading 4"
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="5"]:before,.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="5"]:before {
    content: "Heading 5"
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="6"]:before,.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="6"]:before {
    content: "Heading 6"
}

.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="1"]:before {
    font-size: 2em
}

.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="2"]:before {
    font-size: 1.5em
}

.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="3"]:before {
    font-size: 1.17em
}

.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="4"]:before {
    font-size: 1em
}

.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="5"]:before {
    font-size: .83em
}

.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="6"]:before {
    font-size: .67em
}

.ql-snow .ql-picker.ql-font {
    width: 108px
}

.ql-snow .ql-picker.ql-font .ql-picker-label:before,.ql-snow .ql-picker.ql-font .ql-picker-item:before {
    content: "Sans Serif"
}

.ql-snow .ql-picker.ql-font .ql-picker-label[data-value=serif]:before,.ql-snow .ql-picker.ql-font .ql-picker-item[data-value=serif]:before {
    content: "Serif"
}

.ql-snow .ql-picker.ql-font .ql-picker-label[data-value=monospace]:before,.ql-snow .ql-picker.ql-font .ql-picker-item[data-value=monospace]:before {
    content: "Monospace"
}

.ql-snow .ql-picker.ql-font .ql-picker-item[data-value=serif]:before {
    font-family: Georgia,Times New Roman,serif
}

.ql-snow .ql-picker.ql-font .ql-picker-item[data-value=monospace]:before {
    font-family: Monaco,Courier New,monospace
}

.ql-snow .ql-picker.ql-size {
    width: 98px
}

.ql-snow .ql-picker.ql-size .ql-picker-label:before,.ql-snow .ql-picker.ql-size .ql-picker-item:before {
    content: "Normalx"
}

.ql-snow .ql-picker.ql-size .ql-picker-label[data-value=small]:before,.ql-snow .ql-picker.ql-size .ql-picker-item[data-value=small]:before {
    content: "Small"
}

.ql-snow .ql-picker.ql-size .ql-picker-label[data-value=large]:before,.ql-snow .ql-picker.ql-size .ql-picker-item[data-value=large]:before {
    content: "Large"
}

.ql-snow .ql-picker.ql-size .ql-picker-label[data-value=huge]:before,.ql-snow .ql-picker.ql-size .ql-picker-item[data-value=huge]:before {
    content: "Huge"
}

.ql-snow .ql-picker.ql-size .ql-picker-item[data-value=small]:before {
    font-size: 10px
}

.ql-snow .ql-picker.ql-size .ql-picker-item[data-value=large]:before {
    font-size: 18px
}

.ql-snow .ql-picker.ql-size .ql-picker-item[data-value=huge]:before {
    font-size: 32px
}

.ql-snow .ql-picker.ql-line-height .ql-picker-label {
    width: 28px;
    height: 24px
}

.ql-snow .ql-picker.ql-line-height .ql-picker-item[data-value="1"]:after {
    content: "1"
}

.ql-snow .ql-picker.ql-line-height .ql-picker-item[data-value="1.5"]:after {
    content: "1.5"
}

.ql-snow .ql-picker.ql-line-height .ql-picker-item[data-value="1.6"]:after {
    content: "1.6\ff08\9ed8\8ba4\ff09"
}

.ql-snow .ql-picker.ql-line-height .ql-picker-item[data-value="1.75"]:after {
    content: "1.75"
}

.ql-snow .ql-picker.ql-line-height .ql-picker-item[data-value="2"]:after {
    content: "2"
}

.ql-snow .ql-picker.ql-line-height .ql-picker-item[data-value="2.5"]:after {
    content: "2.5"
}

.ql-snow .ql-picker.ql-line-height .ql-picker-item[data-value="3"]:after {
    content: "3"
}

.ql-snow .ql-picker.ql-line-height .ql-picker-item[data-value="4"]:after {
    content: "4"
}

.ql-snow .ql-color-picker.ql-background .ql-picker-item {
    background-color: #fff
}

.ql-snow .ql-color-picker.ql-color .ql-picker-item {
    background-color: #000
}

.ql-code-block-container {
    position: relative
}

.ql-code-block-container .ql-ui {
    right: 5px;
    top: 5px
}

.ql-toolbar.ql-snow {
    border: 1px solid #ccc;
    box-sizing: border-box;
    font-family: Helvetica Neue,Helvetica,Arial,sans-serif;
    padding: 8px
}

.ql-toolbar.ql-snow .ql-formats {
    margin-right: 15px
}

.ql-toolbar.ql-snow .ql-picker-label {
    border: 1px solid transparent
}

.ql-toolbar.ql-snow .ql-picker-options {
    border: 1px solid transparent;
    box-shadow: #0003 0 2px 8px
}

.ql-toolbar.ql-snow .ql-picker.ql-expanded .ql-picker-label,.ql-toolbar.ql-snow .ql-picker.ql-expanded .ql-picker-options {
    border-color: #ccc
}

.ql-toolbar.ql-snow .ql-color-picker .ql-picker-item.ql-selected,.ql-toolbar.ql-snow .ql-color-picker .ql-picker-item:hover {
    border-color: #000
}

.ql-toolbar.ql-snow+.ql-container.ql-snow {
    border-top: 0px
}

.ql-snow .ql-tooltip {
    background-color: #fff;
    border: 1px solid #ccc;
    box-shadow: 0 0 5px #ddd;
    color: #444;
    padding: 5px 12px;
    white-space: nowrap
}

.ql-snow .ql-tooltip:before {
    content: "Visit URL:";
    line-height: 26px;
    margin-right: 8px
}

.ql-snow .ql-tooltip input[type=text] {
    display: none;
    border: 1px solid #ccc;
    font-size: 13px;
    height: 26px;
    margin: 0;
    padding: 3px 5px;
    width: 170px
}

.ql-snow .ql-tooltip a.ql-preview {
    display: inline-block;
    max-width: 200px;
    overflow-x: hidden;
    text-overflow: ellipsis;
    vertical-align: top
}

.ql-snow .ql-tooltip a.ql-action:after {
    border-right: 1px solid #ccc;
    content: "Edit";
    margin-left: 16px;
    padding-right: 8px
}

.ql-snow .ql-tooltip a.ql-remove:before {
    content: "Remove";
    margin-left: 8px
}

.ql-snow .ql-tooltip a {
    line-height: 26px
}

.ql-snow .ql-tooltip.ql-editing a.ql-preview,.ql-snow .ql-tooltip.ql-editing a.ql-remove {
    display: none
}

.ql-snow .ql-tooltip.ql-editing input[type=text] {
    display: inline-block
}

.ql-snow .ql-tooltip.ql-editing a.ql-action:after {
    border-right: 0px;
    content: "Save";
    padding-right: 0
}

.ql-snow .ql-tooltip[data-mode=link]:before {
    content: "Enter link:"
}

.ql-snow .ql-tooltip[data-mode=formula]:before {
    content: "Enter formula:"
}

.ql-snow .ql-tooltip[data-mode=video]:before {
    content: "Enter video:"
}

.ql-snow a {
    color: #06c
}

.ql-container.ql-snow {
    border: 1px solid #ccc
}

@charset "UTF-8";html,body,div,span,applet,object,iframe,p,blockquote,a,abbr,acronym,address,big,cite,code,del,dfn,img,ins,kbd,q,s,samp,small,strike,sub,sup,tt,var,b,i,center,dl,dt,dd,ol,ul,li,fieldset,form,label,legend,table,caption,tbody,tfoot,thead,tr,th,td,article,aside,canvas,details,embed,figure,figcaption,footer,header,hgroup,menu,nav,output,ruby,section,summary,time,mark,audio,video {
    margin: 0;
    padding: 0;
    border: 0;
    font-size: 100%;
    font: inherit;
    -webkit-tap-highlight-color: transparent
}

::selection {
    color: #fff;
    background: #19988B
}

article,aside,details,figcaption,figure,footer,header,hgroup,menu,nav,section {
    display: block
}

ol,ul {
    list-style: none
}
blockquote {
	border-left: 3px solid #ccc;
	margin: 0;
	padding: 10px;
	background: #eee;
	border-radius: 5px;!important
}

blockquote,q {
    quotes: none
}

blockquote:before,blockquote:after,q:before,q:after {
    content: "";
    content: none
}

table {
    border-collapse: collapse;
    border-spacing: 0
}

a {
    text-decoration: none;
    color: inherit;
    cursor: pointer
}

input[type=number]::-webkit-inner-spin-button,input[type=number]::-webkit-outer-spin-button {
    -webkit-appearance: none
}

img {
    display: block;
    border: none;
 	page-break-inside: avoid; 
	max-width: 100% !important;
}
img.emoji {
	display: inline-block;
   	height: 1em;
	width: 1em;
   	margin: 0 .05em 0 .1em;
   	vertical-align: -0.1em;
}
html {
    width: 100%;
    height: 100%;
    // font-size: 14px;
    font-family: PingFang SC,Tahoma,Helvetica,Arial,Hiragino Sans GB,Microsoft YaHei,Heiti SC,WenQuanYi Micro Hei,sans-serif
}

html body {
    width: 100%;
    height: 100%;
    margin: 0;
	line-height:1.5em;
    overflow-x: hidden;
    box-sizing: border-box
}
.article {
	border-radius: 5px;
	background: #F5F7FA;
}
h1{font-size:28px;line-height:1.2em;display:block;font-weight:bold; !important;}
h2{font-size:24px;line-height:1.2em;display:block;font-weight:bold; !important;}
h3{font-size:18.72px;line-height:1.2em;display:block;font-weight:bold; !important;}
h4{font-size:16px;line-height:1.2em;display:block;font-weight:bold; !important;}
h5{font-size:13.28px;line-height:1.2em;display:block;font-weight:bold; !important;}
h6{font-size:12px;line-height:1.2em;display:block;font-weight:bold; !important;}

html .message-container-overlay {
    position: fixed;
    left: 50%;
    top: 50px;
    width: 400px;
    margin-left: -200px;
    margin-top: 10px;
    z-index: 99
}

textarea {
    font-family: PingFang SC,Tahoma,Helvetica,Arial,Hiragino Sans GB,Microsoft YaHei,Heiti SC,WenQuanYi Micro Hei,sans-serif
}

.ng-relative {
    position: relative
}

.emoji_span {
    height: 24px;
    display: inline-block;
    vertical-align: middle
}

.emoji_span .emoji_local {
    width: 20px;
    height: 20px
}

.fr-quick-insert {
    display: none
}

.fr-view {
    font-family: PingFangSC-Regular!important;
    font-size: 18px!important;
    color: #4d4d4d!important;
    line-height: 28px!important
}

.fr-view p {
    margin: 0;
    padding: 0
}

.fr-view .title-container {
    display: none;
    opacity: 0
}

.fr-toolbar.fr-top {
    border-radius: 0!important;
    border-top: 3px solid #1a1a1a
}

.second-toolbar {
    border-radius: 0!important
}

.file-icon {
    background-image: url(assets/resources/sprite@1x.0a9b898d27588c75.png)
}

.file-icon.file-doc,.file-icon.file-docx {
    background-position: -101px -677px!important
}

.file-icon.file-pdf {
    background-position: -341px -677px!important
}

.file-icon.file-xls,.file-icon.file-xlsx {
    background-position: -221px -677px!important
}

.file-icon.file-ppt,.file-icon.file-pptx,.file-icon.file-pps,.file-icon.file-ppsx {
    background-position: -281px -677px!important
}

.file-icon.file-txt {
    background-position: -161px -677px!important
}

.file-icon.file-mp3,.file-icon.file-wav,.file-icon.file-m4a {
    background-position: -341px -636px!important
}

.file-icon-big {
    background-image: url(assets/resources/sprite@1x.0a9b898d27588c75.png)
}

.file-icon-big.file-doc,.file-icon-big.file-docx {
    background-position: -86px -769px!important
}

.file-icon-big.file-pdf {
    background-position: -486px -769px!important
}

.file-icon-big.file-xls,.file-icon-big.file-xlsx {
    background-position: -286px -769px!important
}

.file-icon-big.file-ppt,.file-icon-big.file-pptx,.file-icon-big.file-pps,.file-icon-big.file-ppsx {
    background-position: -386px -769px!important
}

.file-icon-big.file-txt {
    background-position: -186px -769px!important
}

.file-icon-big.file-mp3,.file-icon-big.file-wav,.file-icon-big.file-m4a {
    background-position: -852px -769px!important
}

.link-of-topic:before {
    background-image: url(assets/resources/sprite@1x.0a9b898d27588c75.png)!important;
    background-position: -387px -575px;
    height: 14px;
    width: 14px;
    content: " ";
    display: inline-block;
    padding-right: 5px;
    margin-left: 4px;
    line-height: 14px
}

.link-of-topic {
    word-break: break-all;
    color: #567895
}

.link-of-topic:hover {
    text-decoration: underline
}

.link-of-topic.noLinkIcon:before {
    display: none
}

@media only screen and (-webkit-min-device-pixel-ratio: 2),only screen and (min--moz-device-pixel-ratio: 2),only screen and (-ms-min-device-pixel-ratio: 2),only screen and (-o-min-device-pixel-ratio: 2),only screen and (min-device-pixel-ratio: 2) {
    .link-of-topic:before {
        background-image:url(assets/resources/sprite@2x.f70e566c4f218699.png)!important;
        background-size: 1024px 1057px
    }

    .file-icon {
        background-image: url(assets/resources/sprite@2x.f70e566c4f218699.png)!important;
        background-size: 1024px 1057px
    }
}

.examination-detail-panel .header-container,.topic-detail-panel .header-container {
    padding: 0!important;
    margin-bottom: 20px
}

.examination-detail-panel .header-container .topic-flag,.topic-detail-panel .header-container .topic-flag {
    top: -30px!important;
    right: -30px!important
}

.examination-detail-panel .avatar,.topic-detail-panel .avatar {
    width: 36px!important;
    height: 36px!important
}

.examination-detail-panel .talk-content-container,.examination-detail-panel .question-content-container,.examination-detail-panel .task-content-container,.examination-detail-panel .solution-content-container,.examination-detail-panel .answer-content-container,.topic-detail-panel .talk-content-container,.topic-detail-panel .question-content-container,.topic-detail-panel .task-content-container,.topic-detail-panel .solution-content-container,.topic-detail-panel .answer-content-container {
    padding: 0!important;
    margin-bottom: 20px!important
}

.examination-detail-panel .question-content-container .join-time,.topic-detail-panel .question-content-container .join-time {
    position: relative;
    top: -5px;
    margin-bottom: 15px!important
}

.examination-detail-panel .talk-content-container .content,.examination-detail-panel .answer-content-container .answer,.topic-detail-panel .talk-content-container .content,.topic-detail-panel .answer-content-container .answer {
    margin-bottom: 20px
}

.examination-detail-panel .comment-item-container,.topic-detail-panel .comment-item-container {
    padding: 0!important
}

.examination-detail-panel .main-comment-item .comment-container,.topic-detail-panel .main-comment-item .comment-container {
    margin-top: 20px!important
}

.examination-detail-panel .reply-comment .container :first-child .comment-item-container,.topic-detail-panel .reply-comment .container :first-child .comment-item-container {
    margin-top: 0!important
}

.examination-detail-panel .reply-comment .container .comment-item-container,.topic-detail-panel .reply-comment .container .comment-item-container {
    width: auto
}

.examination-detail-panel .reply-comment .comment-container,.topic-detail-panel .reply-comment .comment-container {
    margin-top: 12px!important;
    margin-bottom: 0!important;
    padding-right: 0!important
}

.examination-detail-panel .reply-comment .comment-container .emoji-panel,.topic-detail-panel .reply-comment .comment-container .emoji-panel {
    left: 27px!important
}

.examination-detail-panel .comment-item .comment-container,.topic-detail-panel .comment-item .comment-container {
    padding: 0!important
}

.examination-detail-panel .comment-item .comment-container .content-container,.topic-detail-panel .comment-item .comment-container .content-container {
    width: 100%!important;
    background: #fff
}

.examination-detail-panel .icon-comment .comment-container,.topic-detail-panel .icon-comment .comment-container {
    padding: 0!important
}

.examination-detail-panel .icon-comment .comment-container .content-container,.topic-detail-panel .icon-comment .comment-container .content-container {
    width: 100%!important
}

.examination-detail-panel .question-content-container,.topic-detail-panel .question-content-container {
    margin-bottom: 0!important
}

.examination-detail-panel .question-content-container .header,.topic-detail-panel .question-content-container .header {
    margin-bottom: 15px!important;
    font-size: 14px!important
}

.examination-detail-panel .question-content-container .header .fee,.examination-detail-panel .question-content-container .header .owner,.topic-detail-panel .question-content-container .header .fee,.topic-detail-panel .question-content-container .header .owner {
    line-height: 20px!important
}

.examination-detail-panel .question-content-container .answer-question,.topic-detail-panel .question-content-container .answer-question {
    padding-bottom: 0!important
}

.examination-detail-panel .question-content-container .horizontal-line,.topic-detail-panel .question-content-container .horizontal-line {
    margin-right: 0!important
}

.examination-detail-panel .question-content-container .answer-question,.topic-detail-panel .question-content-container .answer-question {
    padding-right: 0!important
}

.examination-detail-panel .operation-icon,.topic-detail-panel .operation-icon {
    margin-top: 20px!important
}

.examination-detail-panel .tag-container,.topic-detail-panel .tag-container {
    padding: 0!important;
    margin-bottom: -10px!important
}

.examination-detail-panel .tag-container .tag,.topic-detail-panel .tag-container .tag {
    margin-bottom: 10px!important
}

.examination-detail-panel .image-gallery-container,.topic-detail-panel .image-gallery-container {
    margin: 0!important
}

.examination-detail-panel .file-gallery-container,.topic-detail-panel .file-gallery-container {
    margin-top: 20px;
    margin-bottom: 20px
}

.examination-detail-panel .video-gallery-container .processing,.topic-detail-panel .video-gallery-container .processing {
    width: 540px!important
}

.examination-detail-panel .video-gallery-container.is-referenced .processing,.topic-detail-panel .video-gallery-container.is-referenced .processing {
    width: 508px!important
}

.examination-detail-panel .referenced-topic,.topic-detail-panel .referenced-topic {
    margin: 0!important
}

#topic-detail-container .skeleton-container .skeleton {
    border-radius: 4px!important
}

.topic-detail-preview .topic-detail-panel .image-gallery-container {
    pointer-events: none
}

.topic-container .talk-content-container .content,.topic-container .answer-content-container .content,.topic-container .solution-content-container .content,.topic-container .task-content-container .content,.topic-container .answer-content-container .answer {
    font-size: 15px!important;
    line-height: 25px!important
}

.topic-detail-panel .talk-content-container .content,.topic-detail-panel .answer-content-container .content,.topic-detail-panel .solution-content-container .content,.topic-detail-panel .task-content-container .content,.topic-detail-panel .question-content-container .content,.topic-detail-panel .talk-content-container .answer,.topic-detail-panel .answer-content-container .answer,.topic-detail-panel .solution-content-container .answer,.topic-detail-panel .task-content-container .answer,.topic-detail-panel .question-content-container .answer {
    font-size: 15px!important;
    line-height: 28px!important
}

.topic-detail-panel .talk-content-container .question,.topic-detail-panel .answer-content-container .question,.topic-detail-panel .solution-content-container .question,.topic-detail-panel .task-content-container .question,.topic-detail-panel .question-content-container .question {
    margin-bottom: 20px!important
}

.topic-preview .image-gallery-container .one-image .image-container {
    width: 130px;
    height: 130px
}

.topic-preview .image-gallery-container .one-image .image-container .single-img {
    width: 100%!important;
    height: 100%!important
}

.sticky-topic-container .link-of-topic:first-child:before,.best-topic-container .link-of-topic:first-child:before {
    position: relative;
    left: 1px;
    top: 2px
}

.notify-wrapper .dynamic .link-of-topic:first-child:before {
    position: relative;
    top: 2px
}

.joined-group-topic-container .view-more-topic {
    padding-top: 20px!important
}

.joined-group-topic-container .topic-preview .main-content .content {
    -webkit-user-select: none;
    user-select: none
}

.string {
    color: green
}

.key {
    color: red
}

.fr-box.fr-basic.fr-top .fr-wrapper {
    box-shadow: none!important
}

.fr-box.fr-basic .fr-element {
    padding: 30px 60px!important
}

.article-container ol,.article-container ul {
    padding-inline-start:10px}

.article-container a {
    color: #567895
}

.article-container a:hover {
    text-decoration: underline
}

.main-content-container .enter-group,.homework-list-container .enter-group {
    display: none!important
}

.main-content-container .topic-detail-panel,.homework-list-container .topic-detail-panel {
    padding-top: 30px!important
}

.go-top-btn {
    display: none;
    cursor: pointer;
    position: fixed;
    bottom: 60px;
    margin-left: 810px;
    width: 40px;
    height: 40px;
    border-radius: 8px;
    background-color: #fff;
    box-shadow: 0 8px 20px #0000000d
}

.go-top-btn .icon {
    width: 18px;
    height: 17px;
    margin: 12px auto;
    background-image: url(assets/resources/sprite@1x.0a9b898d27588c75.png);
    background-position: -265px -311px
}

.go-top-btn:hover {
    background-color: #f5f6f7
}

.topic-container .comment-box .comment-item-container {
    padding: 12px 12px 0;
    margin: 0 20px 0 56px;
    background-color: #f7f9fa
}

.topic-container .comment-box :first-child .comment-item-container {
    border-top-left-radius: 4px;
    border-top-right-radius: 4px
}

.topic-container .comment-box :last-child .comment-item-container {
    padding-bottom: 12px!important;
    border-bottom-left-radius: 4px;
    border-bottom-right-radius: 4px
}

.topic-container .comment-box .comment-container {
    padding: 12px 0 0 12px;
    margin: 0 20px 0 56px;
    background-color: #f7f9fa
}

.group-coupons-container .card-view .flow-loading .text {
    color: #fff
}

.flatpickr-calendar {
    border: 1px solid rgba(0,0,0,.05)!important;
    box-shadow: -5px 6px 12px #0000000d!important;
    -webkit-user-select: none;
    user-select: none
}

.flatpickr-calendar .flatpickr-months {
    margin: 5px 0 12px
}

.flatpickr-calendar .flatpickr-months .flatpickr-prev-month,.flatpickr-calendar .flatpickr-months .flatpickr-next-month {
    top: 5px
}

.flatpickr-calendar .flatpickr-months .numInputWrapper {
    margin-left: 20px
}

#article-preview-container .article-content {
    letter-spacing: 0;
    color: #1a1a1a;
    font-size: 16px;
    word-wrap: break-word
}

#article-preview-container .article-content figure {
    margin-inline-start:0px;margin-inline-end:0px}

#article-preview-container .article-content img {
    object-fit: cover;
    max-width: 100%
}

#article-preview-container .article-content img.fr-dib {
    margin: 5px auto;
    display: block;
    float: none
}

#article-preview-container .article-content img.fr-dib.fr-fil {
    margin-left: 0;
    text-align: left
}

#article-preview-container .article-content img.fr-dib.fr-fir {
    margin-right: 0;
    text-align: right
}

#article-preview-container .article-content a {
    text-decoration: none;
    color: #567895
}

#article-preview-container .article-content a:hover {
    text-decoration: underline
}

#article-preview-container .article-content p {
    margin: 0
}

#article-preview-container .article-content table {
    display: block;
    overflow: auto;
    border: none;
    border-collapse: collapse;
    empty-cells: show;
    max-width: 100%;
    border-spacing: 2px
}

#article-preview-container .article-content table td {
    min-width: 5px
}

#article-preview-container .article-content table th {
    background: #ececec
}

#article-preview-container .article-content code {
    white-space: normal
}

#article-preview-container .article-content pre {
    white-space: pre-wrap
}

#article-preview-container .article-content code {
    background-color: #f0f0f0;
    padding: 2px 4px;
    border-radius: 3px
}

#article-preview-container .article-content .ql-syntax {
    background-color: #23241f;
    color: #f8f8f2;
    overflow: visible;
    white-space: pre-wrap;
    margin-bottom: 5px;
    margin-top: 5px;
    padding: 5px 10px;
    border-radius: 3px
}

#article-preview-container .article-content table td,#article-preview-container .article-content table th {
    border: 1px solid #ddd
}

#article-preview-container .article-content table td:empty,#article-preview-container .article-content table th:empty {
    height: 20px
}

.examination-list-container .examination-list-main .image-gallery-container {
    margin-bottom: 0
}

.examination-list-container .examination-list-main .one-image .image-container {
    width: 130px;
    height: 130px
}

.examination-list-container .examination-list-main .one-image .image-container .single-img {
    width: 100%!important;
    height: 100%!important
}

.video-container .video-js {
    font-size: 14px
}

.video-container .video-js .vjs-control-bar,.video-container .video-js .vjs-volume-vertical {
    background-color: #0003
}

.video-container .video-js .vjs-time-control {
    display: block
}

.video-container .video-js .vjs-remaining-time {
    display: none
}

.video-container .video-js.vjs-16-9,.video-container .video-js.vjs-4-3 {
    height: 504px
}

.video-container .video-js .vjs-tech,.video-container .video-js button {
    outline: none
}

.video-container .video-js .vjs-big-play-button {
    height: 50px;
    width: 50px;
    border-radius: 50%;
    margin-top: -25px;
    margin-left: -25px;
    border-width: 1px
}

.video-container .video-js .vjs-volume-level:before {
    left: -.5em
}

.video-container .video-js .vjs-play-progress:before {
    top: -.4em
}

.video-container .video-js .vjs-play-progress,.video-container .video-js .vjs-volume-level {
    background-color: #19988b
}

.video-container .video-js.vjs-paused .vjs-big-play-button {
    display: block
}

.video-container .video-js:hover .vjs-big-play-button {
    background-color: #00000080
}

.ql-snow .ql-editor .ql-code-block-container {
    background-color: #00000008;
    border: 1px solid #f0f0f0;
    border-radius: 2px;
    color: #333
}

.ql-snow .ql-editor blockquote {
    color: #9a9a9a;
    margin: 15px 0;
    border-left: 3px solid #ccc;
    padding-left: 10px
}

.ql-editor li>.ql-ui:before {
    text-align: center
}

@media only screen and (-webkit-min-device-pixel-ratio: 2),only screen and (min--moz-device-pixel-ratio: 2),only screen and (-ms-min-device-pixel-ratio: 2),only screen and (-o-min-device-pixel-ratio: 2),only screen and (min-device-pixel-ratio: 2) {
    .go-top-btn .icon {
        background-image:url(assets/resources/sprite@2x.f70e566c4f218699.png)!important;
        background-size: 1024px 1057px
    }
}

@media screen and (max-width: 1600px) {
    .go-top-btn {
        margin-left:610px
    }
}
	</style>
<script src="https://twemoji.maxcdn.com/v/latest/twemoji.min.js"></script>
<script>window.onload = function () { twemoji.parse(document.body);}</script>
</head>
<body>
`
	return
}

func GenEndHtml() (result string) {
	result += `
</body>
</html>`
	return
}

func GenPageBreak() (result string) {
	result = `<p style="page-break-after: always;">`
	return
}

func GenHLevelHtml(level int, startTag bool) (result string) {
	sTag := map[int]string{0: `<h1>`, 1: `<h2>`, 2: `<h3>`, 3: `<h4>`, 4: `<h5>`, 5: `<h6>`}
	eTag := map[int]string{0: `</h1>`, 1: `</h2>`, 2: `</h3>`, 3: `</h4>`, 4: `</h5>`, 5: `</h6>`}
	if startTag {
		if tag, ok := sTag[level]; ok {
			result = tag
		}
	} else {
		if tag, ok := eTag[level]; ok {
			result = tag
		}
	}
	return
}
