// Copyright (c) 2022 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package pmapi

func iHasFlag(i, flag int) bool           { return i&flag == flag }
func iHasAtLeastOneFlag(i, flag int) bool { return i&flag > 0 }
func iIsFlag(i, flag int) bool            { return i == flag }
func iHasNoneOfFlag(i, flag int) bool     { return !iHasAtLeastOneFlag(i, flag) }
