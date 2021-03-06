/*
 * Copyright (C) 2019 ~ 2020 Uniontech Software Technology Co.,Ltd
 *
 * Author:
 *
 * Maintainer:
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package xmlp

import (
	"bytes"
	"strings"
)

// snake_case to CamelCase
func snake2Camel(name string) string {
	//name = strings.ToLower(name)
	var out bytes.Buffer
	for _, word := range strings.Split(name, "_") {
		word = strings.ToLower(word)
		//if subst, ok := config.word_subst[word]; ok {
		//out.WriteString(subst)
		//continue
		//}

		if word == "" {
			out.WriteString("_")
			continue
		}
		out.WriteString(strings.ToUpper(word[0:1]))
		out.WriteString(word[1:])
	}
	return out.String()
}

// name is c identifier
func (oi *ObjectInfo) GetFunctionInfo(name string) *FunctionInfo {
	for _, funcInfoList := range [][]*FunctionInfo{
		oi.Functions,
		oi.Constructors,
		oi.Methods,
	} {
		funcInfo := getFuncByName(funcInfoList, name)
		if funcInfo != nil {
			return funcInfo
		}
	}
	return nil
}

func (si *StructInfo) GetFunctionInfo(name string) *FunctionInfo {
	for _, funcInfoList := range [][]*FunctionInfo{
		si.Functions,
		si.Constructors,
		si.Methods,
	} {
		funcInfo := getFuncByName(funcInfoList, name)
		if funcInfo != nil {
			return funcInfo
		}
	}
	return nil
}

func (ii *InterfaceInfo) GetFunctionInfo(name string) *FunctionInfo {
	for _, funcInfoList := range [][]*FunctionInfo{
		ii.Functions,
		ii.Methods,
	} {
		funcInfo := getFuncByName(funcInfoList, name)
		if funcInfo != nil {
			return funcInfo
		}
	}
	return nil
}

func getFuncByName(funcInfoList []*FunctionInfo, name string) *FunctionInfo {
	for _, funcInfo := range funcInfoList {
		if funcInfo.CIdentifier == name {
			return funcInfo
		}
	}
	return nil
}

func (ns *Namespace) GetFunctionInfo(name string) *FunctionInfo {
	for _, objInfo := range ns.Objects {
		funcInfo := objInfo.GetFunctionInfo(name)
		if funcInfo != nil {
			return funcInfo
		}
	}

	for _, structInfo := range ns.Structs {
		funcInfo := structInfo.GetFunctionInfo(name)
		if funcInfo != nil {
			return funcInfo
		}
	}

	for _, ifcInfo := range ns.Interfaces {
		funcInfo := ifcInfo.GetFunctionInfo(name)
		if funcInfo != nil {
			return funcInfo
		}
	}

	return nil
}

type IGetFunctionInfo interface {
	GetFunctionInfo(name string) *FunctionInfo
}
