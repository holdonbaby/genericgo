/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   interface.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Xu Huang <xuh15462@gmail.com>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/05/01 22:26:34 by 黄旭                #+#    #+#             */
/*   Updated: 2024/05/01 22:28:22 by Xu Huang         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// file header only in the interface file.go
package utils

import "fmt"

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

type Comparable interface {
	Number | ~string
}

type MapComparable interface {
	Comparable
	fmt.Stringer
}
